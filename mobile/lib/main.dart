import 'package:flutter/material.dart';
import 'package:toclido/screens/signup.dart';

void main() {
  runApp(const AppRoot());
}

class AppRoot extends StatelessWidget {
  const AppRoot({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: SignupScreen(),
    );
  }
}
