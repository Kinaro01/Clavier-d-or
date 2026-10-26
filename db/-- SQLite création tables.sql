-- SQLite

CREATE TABLE questions (
    id INTEGER PRIMARY KEY,
    theme_id INTEGER NOT NULL,
    type_question VARCHAR(50) NOT NULL, -- 'normal' ou 'boss'
    libelle TEXT NOT NULL,
    option_a TEXT NOT NULL,
    option_b TEXT NOT NULL,
    option_c TEXT NOT NULL,
    option_d TEXT NOT NULL,
    est_correcte CHAR(1) NOT NULL,
    indice TEXT NOT NULL,
    FOREIGN KEY (theme_id) REFERENCES theme(id)
)


CREATE TABLE theme (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nom_theme TEXT NOT NULL
)


CREATE TABLE sauvegarde (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    pseudo TEXT NOT NULL,
    date TEXT DEFAULT (datetime('now')),
    score_total INTEGER
)

CREATE TABLE sauvegardeTheme (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sauvegarde_id INTEGER NOT NULL,
    theme_id INTEGER NOT NULL,
    theme_demarre BOOL,
    FOREIGN KEY (sauvegarde_id) REFERENCES sauvegarde(id),
    FOREIGN KEY (theme_id) REFERENCES theme(id)
)

CREATE TABLE sauvegardeQuestion (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    sauvegarde_id_theme INTEGER NOT NULL,
    sauvegarde_id INTEGER NOT NULL,
    question_id INTEGER NOT NULL,
    complete BOOL,
    score_question INTEGER,
    FOREIGN KEY (sauvegarde_id_theme) REFERENCES sauvegardeTheme(id),
    FOREIGN KEY (sauvegarde_id) REFERENCES sauvegarde(id)
    FOREIGN KEY (question_id) REFERENCES questions(id)
)

