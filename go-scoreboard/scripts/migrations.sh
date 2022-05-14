#!/bin/sh

sqlite3 ~/scoreboard.db <<EOF
  CREATE TABLE snookers (
      id     INTEGER PRIMARY KEY,
      winner TEXT NOT NULL,
      loser  TEXT NOT NULL,
      diff   INTEGER NOT NULL,
      date   TEXT DEFAULT (datetime('now'))
  );

  CREATE TABLE dees (
      id          INTEGER PRIMARY KEY,
      winner      TEXT NOT NULL,
      loser1      TEXT NOT NULL,
      loser1_card INTEGER NOT NULL,
      loser2      TEXT NOT NULL,
      loser2_card INTEGER NOT NULL,
      loser3      TEXT NOT NULL,
      loser3_card INTEGER NOT NULL,
      date        TEXT DEFAULT (datetime('now'))
  );

  CREATE TABLE landlords (
      id       INTEGER PRIMARY KEY,
      landlord TEXT NOT NULL,
      farmer1  TEXT NOT NULL,
      farmer2  TEXT NOT NULL,
      win      INTEGER NOT NULL,
      date     TEXT DEFAULT (datetime('now'))
  );
EOF
