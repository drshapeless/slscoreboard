import 'package:flutter/material.dart';
import 'package:flutter_scoreboard/snookers.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Shit'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  final TextEditingController _usernameController1 = TextEditingController();
  final TextEditingController _usernameController2 = TextEditingController();

  final TextEditingController _passController = TextEditingController();

  Future<Snooker>? _futureSnooker;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            _buildPassTextField(),
            // _buildUsernameRow(),
          ],
        ),
      ),
    );
  }

  TextField _buildPassTextField() {
    return TextField(
      controller: _passController,
      decoration: const InputDecoration(
        hintText: "pass",
        hintStyle: TextStyle(
          color: Colors.grey,
        ),
      ),
    );
  }

  Row _buildUsernameRow() {
    return Row(
      children: [
        TextField(
          controller: _usernameController1,
          decoration: const InputDecoration(
            hintText: "Player 1",
            hintStyle: TextStyle(
              color: Colors.grey,
            ),
          ),
        ),
        TextField(
          controller: _usernameController2,
          decoration: const InputDecoration(
            hintText: "Player 2",
            hintStyle: TextStyle(
              color: Colors.grey,
            ),
          ),
        ),
      ],
    );
  }
}
