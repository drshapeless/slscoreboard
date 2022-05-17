import 'package:flutter/material.dart';
import 'package:flutter_scoreboard/snookers.dart';

void main() {
  runApp(const SnookerUploadApp());
}

class SnookerUploadApp extends StatelessWidget {
  const SnookerUploadApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Snooker Upload",
      theme: ThemeData.dark().copyWith(),
      home: const SnookerUploadHome(title: "Snooker home"),
    );
  }
}

class SnookerUploadHome extends StatefulWidget {
  const SnookerUploadHome({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<SnookerUploadHome> createState() => _SnookerUploadHomeState();
}

class _SnookerUploadHomeState extends State<SnookerUploadHome> {
  TextEditingController _usernameController1 = TextEditingController();
  TextEditingController _usernameController2 = TextEditingController();
  TextEditingController _passController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // backgroundColor: Colors.black,
      body: Column(
        children: [
          TextField(
            controller: _passController,
            decoration: InputDecoration(
              labelText: "Pass",
              prefixIcon: Icon(Icons.lock),
            ),
          ),
          _buildPlayerTextfields(),
          Padding(
            padding: EdgeInsets.only(bottom: 10),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              _build8balls(),
              _build8balls(),
            ],
          ),
          Divider(
            color: Colors.white,
          ),
        ],
      ),
    );
  }

  Row _buildPlayerTextfields() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Expanded(
          child: _buildPlayerTextfield1(_usernameController1, "Player 1"),
        ),
        Expanded(
          child: _buildPlayerTextfield1(_usernameController2, "Player 2"),
        ),
      ],
    );
  }

  TextField _buildPlayerTextfield1(TextEditingController con, String label) {
    return TextField(
      controller: con,
      decoration: InputDecoration(
        labelText: label,
        prefixIcon: const Icon(Icons.person),
      ),
    );
  }

  Column _build8balls() {
    double padding = 20;
    return Column(
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            _buildBallButtonWithPadding("0", padding),
            _buildBallButtonWithPadding("1", padding),
            _buildBallButtonWithPadding("2", padding),
            _buildBallButtonWithPadding("3", padding),
          ],
        ),
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            _buildBallButtonWithPadding("4", padding),
            _buildBallButtonWithPadding("5", padding),
            _buildBallButtonWithPadding("6", padding),
            _buildBallButtonWithPadding("7", padding),
          ],
        )
      ],
    );
  }

  Padding _buildBallButtonWithPadding(String label, double padding) {
    return Padding(
      padding: EdgeInsets.all(padding),
      child: _buildBallButton1(label),
    );
  }

  MaterialColor c = Colors.yellow;

  void changeColor(bool t) {
    if (t) {
      setState(() {
        c = Colors.red;
    });
    } else {
      setState(() {
        c = Colors.yellow;
    });
    }
  }

  ElevatedButton _buildBallButton1(String label) {
    return ElevatedButton(
      onPressed: () {},
      onHover: changeColor,
      child: Text(
        label,
        style: TextStyle(
          color: Colors.black,
          fontSize: 20,
          fontWeight: FontWeight.bold,
        ),
      ),
      style: ElevatedButton.styleFrom(
        primary: c,
        shape: CircleBorder(),
        padding: EdgeInsets.all(20),
      ),
    );
  }
}
