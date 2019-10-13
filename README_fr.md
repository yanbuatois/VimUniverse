#VimUniverse
## Histoire
En vous aventurant aux confins du territoire de Xinul, vous atterrissez dans une jungle, la légendaire jungle de Vim,
connue par le fait que les nombreux explorateurs qui s'y sont aventurés n'en sont jamais ressorti. Votre tâche sera,
pour survivre, de braver les pièges de cette jungle, en trouvant les objets laissés par les autres explorateurs, et en
affrontant de grands dangers.

## Installation
Téléchargez l'archive ZIP correspondant à votre système d'exploitation et votre architecture. Extrayez l'intégralité de
son contenu dans le même dossier, sans rien renommer, puis lancez le jeu dans le terminal (utilisez pour cela les fichiers
`run.bat` ou `run.sh` mis à disposition pour vous en assurer).

### Note
Il faut bénéficier d'un terminal ayant un jeu de caractères étendu, ce qui n'est pas le cas de celui de Windows par défaut.
Nous vous recommandons donc l'installation d'un terminal tiers prenant correctement en charge les caractères utilisés par le jeu,
tel que [La Preview de Windows Terminal](https://www.microsoft.com/store/apps/9n0dx20hk701?cid=storebadge&ocid=badge).

## Comment jouer ?
Le jeu est très simple. Il utilise le fonctionnement de l'éditeur de texte VIM. Ainsi, pour vous déplacer, il faut utiliser
les touches `h`, `j`, `k` et `l` de votre clavier. Elles vous permettront, respectivement, d'aller à gauche, en bas, en haut
et à droite. 

La plupart des autres actions, y compris quitter le jeu, nécessite de taper une commande. Les commandes commencent toutes
par le caractère `:`. Taper celui-ci fera directement apparaître la barre de commandes. Pour la faire disparaître, effacez
simplement chacun des caractères que vous y avez tapé. Vous trouverez la liste des commandes et leur effet via la commande
`:help`.

Voici la description des éléments que vous pourrez rencontrer dans ce monde, pour savoir comment les affronter :
* `옷` Le joueur
* `⌴` Le bateau, requis pour traverser des étendues d'eau. Vous devrez le ramasser afin d'explorer tout le jeu. C'est aussi
votre apparence quand vous l'utilisez.
* `~` De l'eau, vous pouvez la traverser dès que vous avez un bateau dans votre inventaire.
* Les blocs blanc : Des arbres épais infranchissables.
* `K` La clé pour s'enfoncer dans la jungle en utilisant les portes que vous trouverez
* `H` Une porte ouverte.
* `ℍ` Une porte verrouillée, nécessitant une clé.

Bonne chance pour vous échapper.

## Créer vos propres niveaux
Afin de créer vos propres niveaux, il vous suffit d'ajouter un fichier ou de modifier ceux présents dans le dossier `levels`.
Le nom correspond au numéro du niveau (ils doivent commencer à 1 et être ininterrompus). Le fichier est un simple fichier
texte dans lequel vous pourrez placer différents éléments sans problème. Voici les caractères admins et leur signification :
* ` ` (espace) Espace libre
* `S` Point d'apparition du joueur
* `B` Bateau
* `W` Arbres épais
* `~` Étendue d'eau
* `K` Clé
* `L` Porte verrouillée
* `O` Porte ouverte

Lorsque le joueur réussit le dernier niveau disponible, il a gagné.

## Évolutions futures
Le jeu devrait évoluer, avec l'ajout d'ennemis, exploitant le système de vies déjà implémenté. Ceux-ci auront pour but de
tuer le joueur, ou de le retenir dans Vim. Ils pourront lâcher des objets comme des clés ou un bateau en mourrant.

Une génération procédurale de l'univers pourrait également être ajoutée, afin de créer des niveaux labyrinthiques différents
à chaque partie de façon automatique.

## Crédits
Jeu réalisé intégralement par [Yan Buatois](https://github.com/yanbuatois).

Utilisation du framework [Termloop](https://github.com/JoelOtter/termloop) par [JoelOtter](https://github.com/JoelOtter)