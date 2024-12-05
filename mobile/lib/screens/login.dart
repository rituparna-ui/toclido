import 'package:flutter/material.dart';
import 'package:toclido/screens/signup.dart';
import 'package:toclido/widgets/login_form.dart';

class LoginScreen extends StatelessWidget {
  const LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('TOCLIDO'),
      ),
      body: SafeArea(
        child: Column(
          children: [
            const LoginForm(),
            GestureDetector(
              onTap: () {
                Navigator.pushReplacement(
                  context,
                  MaterialPageRoute(
                    builder: (context) => const SignupScreen(),
                  ),
                );
              },
              child: Container(
                decoration: const BoxDecoration(
                  color: Colors.teal,
                ),
                width: double.infinity,
                child: const Padding(
                  padding: EdgeInsets.all(16.0),
                  child: Center(
                    child: Text(
                      'Don\'t have an account? Signup',
                      style: TextStyle(
                        fontSize: 20,
                        color: Colors.white,
                      ),
                    ),
                  ),
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
