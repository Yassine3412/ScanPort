# Port Scanner en Go

## Description

Ce projet est un scanner de ports rapide écrit en Go. Il permet de scanner les 65 535 ports d'une adresse IP donnée pour identifier les ports ouverts. Il utilise des **goroutines** pour exécuter les scans en parallèle et un **pool de workers** pour optimiser les performances.

## Fonctionnalités

✅ Scan rapide des ports grâce aux goroutines
✅ Utilisation d'un pool de workers pour éviter la surcharge
✅ Détection des ports ouverts
✅ Temps d'attente configurable pour chaque connexion

## Installation

### 1. Cloner le projet

```sh
git clone https://github.com/Yassine3412/port-scanner.git
cd port-scanner
```

### 2. Installer les dépendances

Assurez-vous d'avoir Go installé, puis exécutez :

```sh
go mod init portscanner
go get github.com/fatih/color
```

## Utilisation

### Lancer un scan sur une adresse IP spécifique

```sh
go run main.go -host 192.168.1.1
```

Par défaut, l'adresse IP scannée est `0.0.0.0`.

### Options disponibles

| Option  | Description                 |
| ------- | --------------------------- |
| `-host` | Spécifie l'adresse IP cible |

## Explication du Code

1. **Initialisation de l'adresse IP** via `flag.String("host", "0.0.0.0", "Adresse IP")`
2. **Utilisation d'un canal (`ports`)** pour gérer la liste des ports à scanner
3. **Création de 100 workers** qui prennent les ports dans la file d'attente et les scannent
4. **Affichage des ports ouverts** en vert grâce à la bibliothèque `github.com/fatih/color`

## Améliorations possibles

- Ajouter un affichage plus détaillé (temps d'exécution, nombre de ports ouverts...)
- Permettre de scanner une plage de ports définie par l'utilisateur
- Exporter les résultats dans un fichier JSON ou CSV

🔥 **Développé avec Go pour une performance maximale !**
