# Projet

Ce script permet de scanner des ports sur une machine spécifier. Ce script a été crée à l'aide de COBRA. 
COBRA est bibliothèque fournissant une interface simple pour créer des interfaces CLI modernes et puissantes,
similaires aux outils git et go.

# Installation

afin de pouvoir utiliser ce script, il suffit de taper la commande suivante

```git clone https://github.com/ArnodTOTO/essir22-prog-go.git```

# Usage

$ cd essir22-prog-go/MyProject

$ ./MyProject scan --help

allows to scan the ports target

Usage:

  MyProject scan [flags]

Flags:

  -h, --help             help for scan

  -p, --port string      design the port target

  -q, --quiet            do not log, only display the results

  -t, --target string    design the target

  -w, --workers string   indicates the number of workers (default "10")

# Exemple

$ ./MyProject scan -p 80 -t localhost 

Starting scan localhost at 2023-02-04 12:12:22.013862054 -0500 EST m=+0.000616858

Port: 80 Open

$./MyProject scan -p all -t localhost -w 16 -q

Starting scan localhost at 2023-02-04 12:14:08.99886702 -0500 EST m=+0.000478498

Port: 21 Open

Port: 22 Open

Port: 80 Open

$./MyProject scan -p 1-80 -t localhost -w 16

INFO[0000] Maybe close  

INFO[0000] Maybe close  

INFO[0000] Maybe close   

INFO[0000] Maybe close 

INFO[0000] Maybe close 

INFO[0000] Maybe close   

INFO[0000] Maybe close   

Open port: 80

INFO[0000] Maybe close  

INFO[0000] Maybe close  

# Note

Ce script doit être utilisé dans un cadre légal afin d'éviter tout problème au niveau de la loi.