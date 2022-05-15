import 'dart:convert';

import 'package:http/http.dart' as http;

import 'globals.dart' as globals;

class Snooker {
  final String winner;
  final String loser;
  final int diff;
  final String date;
  final int id;

  const Snooker({
    required this.winner,
    required this.loser,
    required this.diff,
    required this.date,
    required this.id,
  });

  const Snooker.simple(String winner, String loser, int diff)
      : this(
          winner: winner,
          loser: loser,
          diff: diff,
          date: "",
          id: 0,
        );

  factory Snooker.fromJson(Map<String, dynamic> json) {
    return Snooker(
      winner: json['winner'],
      loser: json['loser'],
      diff: json['diff'],
      date: json['date'],
      id: json['id'],
    );
  }
}

Future<Snooker> createSnooker(String winner, String loser, int diff) async {
  final response = await http.post(
    Uri.parse('${globals.apiBaseURL}/snooker/'),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(<String, dynamic>{
      'pass': globals.apiPass,
      'winner': winner,
      'loser': loser,
      'diff': diff,
    }),
  );

  if (response.statusCode == 201) {
    return Snooker.fromJson(jsonDecode(response.body)['snooker']);
  } else {
    throw Exception("Failed to create snooker.");
  }
}

Future<List<Snooker>> getSnookers(int page) async {
  final response =
      await http.get(Uri.parse("${globals.apiBaseURL}/snooker/$page"));

  if (response.statusCode == 200) {
    List<dynamic> lss = json.decode(response.body)['snookers'];
    List<Snooker> lsno = [];
    for (dynamic i in lss) {
      lsno.add(Snooker.fromJson(i));
    }
    return lsno;
  } else {
    throw Exception("Failed to get snookers.");
  }
}
