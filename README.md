# TP1 de Cryptologie : Rainbow Table

## Build

Le projet est écrit en Go `1.21.5`. Un Makefile est fourni pour compiler le projet.

La première fois lancez `make download` pour télécharger les dépendances.
Ensuite lancez `make` pour compiler le projet, le binaire sera dans le dossier `bin`.

Il existe une version sequentielle et une version parallèle (par défaut). Pour compiler la version séquentielle, utilisez `make seq`.

## LFS

Deux tables arc-en-ciel sont fournies : `table_a26A_s4_500x20000.gob` et `table_a40_s5_1000x1000000.gob`. Elles sont stockées avec Git LFS.

Assurez-vous d'avoir [Git LFS](https://git-lfs.com/) installé et configuré sur votre machine avant de cloner le projet. Sinon, lancez `git lfs pull` après avoir cloné le projet.

## Quickstart

```bash
make
./bin/rbt -h # toutes les sous-commandes ont une aide
./bin/rbt -A 26A -s 4 stats 500 20000 # calcul le coverage pour les paramètres donnés
./bin/rbt -A 26A -s 4 create 500 20000 # si aucun fichier n'est donné, un nom par défaut est généré
./bin/rbt test hash BEEF # 9DBA3007DE04696303B91C7A87554A9BBC62FCE4
./bin/rbt crack 9DBA3007DE04696303B91C7A87554A9BBC62FCE4 ./table_a26A_s4_500x20000.gob # BEEF
./bin/rbt info --none ./table_a26A_s4_500x20000.gob # affiche les informations de la table, mais pas son contenu (incompatible avec --all et --max)
```

## Question 5


*Quelle est la complexité (en temps et en espace) de recherche dans une telle table si la table initiale contenait hauteur lignes et largeur colonnes ?*

SKIPPED

*Comparez cela avec les complexités (en temps et en espace) de la recherche exhaustive et celle du précalcul complet ?*

SKIPPED

## Question 8

*En quoi est-ce que l'ajout du paramètre t dans la fonction h2i permet d'augmenter la couverture de la table ?*

Le paramètre t permet de limiter les collisions. Si nous avons moins de collisions, nous augmentons le nombre de valeurs différentes et donc la converture de la table.

## Question 12

*Estimez la complexité de la recherche dans une table arc-en-ciel.*

## Question 14

Je vais exposer ici les résultats de mon programme, suivant deux version : une version séquentiel, et une parallèle expérimentale tirant partie des goroutines, une fonctionnalité qui s'apparent à des thread légés.

**Pour l'instant, seul la création de la table est parallélisé**

La version séquentiel est compilable avec le tag `seq`.

| _Séquentiel_                 | A=26A & S=4  | A=40 & S=5        |
| ---------------------------- | ------------ | ----------------- |
| largeur x hauteur            | 500 x 20 000 | 1 000 x 1 000 000 |
| couverture                   | 99.33%       | 96.49%            |
| taille de la table           | 171 Ko       | 10 457 Ko         |
| temps de calcul de la table  | ~5.517s      | ~692.136s         |
| temps de calcul de l'inverse | ~0.019s      | 0.391s            |


| _Parallèle_                 | A=26A & S=4 | A=40 & S=5 |
| --------------------------- | ----------- | ---------- |
| temps de calcul de la table | ~0.844s     | ~95.375s   |

`16de25af888480da1af57a71855f3e8c515dcb61 => CODE`
`dafaa5e15a30ecd52c2d1dc6d1a3d8a0633e67e2 => n00b.`

## Question 15

## Question 16

Par recherche exhaustive :

| _Séquentiel_                 | A=26A & S=4 | A=40 & S=5      |
| ---------------------------- | ----------- | --------------- |
| largeur x hauteur            | 2 x 456 976 | 2 x 115 856 201 |
| couverture                   | 100%        | 100%            |
| taille de la table           | 3 889 Ko    | 1 211 655 Ko    |
| temps de calcul de la table  | ~0.135s     | ~230.076s       |
| temps de calcul de l'inverse | ~0.117s     | ~28.707s        |

Le temps de recherche est plus long que par une recherche hybride, mais le résultat est biaisé par le temps que prend la table pour être chargée en mémoire.

Le temps de calcul de la table est lui plus court, car mois de chaînes sont calculées.

La taille de la table est inévitablement plus grande, car il faut stocker toutes les combinaisons possibles.

## Question 17

Le sel est une chaine de caractère ajouté au mot de passe qui rend impossible la création d'une table arc-en-ciel, car il faudrait créer une table par sel (si tenté qu'on le connaisse).

## Pistes d'amélioration

- [ ] Paralléliser la recherche dans la table
- [X] Parallélisation plus intéligente que 1 goroutine par boucle (pool de goroutine/worker ?)
- [ ] Serialiser la table "fait maison" pour ne pas la charger en mémoire pour des commandes simples (stats, info)