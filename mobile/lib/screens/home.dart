import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:toclido/constants.dart';
import 'package:toclido/models/todo.dart';
import 'package:toclido/screens/login.dart';
import 'package:http/http.dart' as http;

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  List<Todo> todosList = [];

  // @override
  // void didChangeDependencies() {
  //   super.didChangeDependencies();
  //   fetchTodos();
  // }

  Future<void> _logout() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    await prefs.clear();
    if (mounted) {
      Navigator.of(context).pushReplacement(
        MaterialPageRoute(
          builder: (context) => const LoginScreen(),
        ),
      );
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('Logged out'),
        ),
      );
    }
  }

  void handleClick(String value) async {
    switch (value) {
      case 'Logout':
        await _logout();
        break;
    }
  }

  Future<void> fetchTodos() async {
    final SharedPreferences prefs = await SharedPreferences.getInstance();
    final String token = prefs.getString('token')!;
    try {
      final response = await http.get(
        Uri.parse('$apiUrl/todos'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );
      final body = jsonDecode(response.body);
      if (response.statusCode == 200) {
        List<Todo> todos = [];
        for (var todo in body['todos']) {
          todos.add(Todo.fromJson(todo));
        }
        setState(() {
          todosList = todos;
        });
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(e.toString()),
          ),
        );
      }
    }
  }

  @override
  void initState() {
    super.initState();
    fetchTodos();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('TOCLIDO'),
        actions: <Widget>[
          PopupMenuButton<String>(
            onSelected: handleClick,
            itemBuilder: (BuildContext context) {
              return {'Logout'}.map((String choice) {
                return PopupMenuItem<String>(
                  value: choice,
                  child: Text(choice),
                );
              }).toList();
            },
          ),
        ],
      ),
      body: ListView.builder(
          itemBuilder: (context, index) {
            return Dismissible(
              key: Key(todosList[index].sId!),
              child: ListTile(
                title: Text(todosList[index].title!),
                subtitle: Text(todosList[index].description!),
                leading: todosList[index].completed!
                    ? GestureDetector(
                        child: const Icon(Icons.check_circle),
                        onTap: () {
                          setState(() {
                            todosList[index].completed = false;
                          });
                        },
                      )
                    : GestureDetector(
                        child: const Icon(Icons.radio_button_unchecked),
                        onTap: () {
                          setState(() {
                            todosList[index].completed = true;
                          });
                        },
                      ),
              ),
            );
          },
          itemCount: todosList.length),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          print("add");
        },
        child: const Icon(Icons.add),
      ),
    );
  }
}
