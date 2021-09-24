![forthebadge](https://forthebadge.com/images/badges/0-percent-optimized.svg) ![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg) ![shield](https://img.shields.io/badge/Usage-serveur-lightgrey) ![shield](https://img.shields.io/badge/Status-Production-green) 

# Proxy-Go V2
## Notice d'emploi

*- Les 2 ports (source et destination) doivent être libre*

**Sur Linux, si l'un des ports est compris entre 0 et 1024, vous devez autoriser le script à utiliser ces ports :**

`sudo setcap CAP_NET_BIND_SERVICE=+eip /path/to/proxygo`

## Options disponibles

-h : help.

-l : port d'écoute (par défaut 80).

-r : port de redirection (par défaut 443).


## Usage

### Sous Linux :

`chmod +x proxygo`

`./proxygo` permet de rediriger le trafic depuis et vers les ports par défaut (80 -> 443).

Exemples :

`./proxygo -l 8080 -r 443` permet de rediriger le trafic du port 8080 vers le 443.

### Sous Windows :

Double cliquer sur `proxygo.exe` pour une rediction depuis et vers les ports par défaut.

Depuis une invite de commande ou PowerShell `proxygo.exe -l 8080 -r 443` pour rediriger le trafic du port 8080 vers le 443.
