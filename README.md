# TP1 de Cryptologie : Rainbow Table

## Build

Le projet est écrit en Go `1.21.5`. Un Makefile est fourni pour compiler le projet.

La première fois lancez `make download` pour télécharger les dépendances.
Ensuite lancez `make` pour compiler le projet, le binaire sera dans le dossier `bin`.

## Question 5


*Quelle est la complexité (en temps et en espace) de recherche dans une telle table si la table initiale contenait hauteur lignes et largeur colonnes ?*

SKIPPED

*Comparez cela avec les complexités (en temps et en espace) de la recherche exhaustive et celle du précalcul complet ?*

SKIPPED

## Question 8

*En quoi est-ce que l'ajout du paramètre t dans la fonction h2i permet d'augmenter la couverture de la table ?*

Le paramètre t permet de limiter les collisions. Si nous avons moins de collisions, nous augmentons le nombre de valeurs différentes et donc la converture de la table.