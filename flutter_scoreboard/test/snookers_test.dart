import 'dart:convert';

import 'package:test/test.dart';
import 'package:flutter_scoreboard/snookers.dart';
import 'package:http/http.dart' as http;

import 'package:flutter_scoreboard/globals.dart' as globals;

void main() {
  String winner = "Jacky";
  String loser = "Eric";
  int diff = 88;

  test("Testing", () async {
    // Snooker s = await createSnooker(winner, loser, diff);
    // print(s);
    List<Snooker> l = await getSnookers(1);
    for (Snooker i in l) {
      print(i.diff);
    }
});
}
