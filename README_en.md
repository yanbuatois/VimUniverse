#VimUniverse
## History
While exploring the Xinul territory, you're lost in the legendary Vim Jungle, well known because of the numerous
explorers who didn't manage to find how to get out. In this game, you will need to go out, foil traps and fight, to find
the exit. You can use items explorers left behind them.

## Installation
Download the ZIP archive, corresponding to your OS and its architecture. Extract its whole content to the same directory,
without renaming anything. Then launch the game inside the terminal (you can use `run.bat` or `run.sh` to ensure it).

### Note
You need to have a Terminal which supports extended charset, which is not the Windows one case, by default.
So we suggest you to use a third-party terminal, correctly handling characters used by the game. On Windows, you can use
the [Windows Terminal Preview](https://www.microsoft.com/store/apps/9n0dx20hk701?cid=storebadge&ocid=badge).id=badge).

## How to play ?
The game is very simple to use. It has the same way to use it as Vim, the popular Linux Text Editor. So to move on the map,
you should use the `h`, `j`, `k` and `l` keys, which respectively make the player move left, bottom, top and right.

Most of other actions, like quitting the game, requires to type a command. Every command starts with the `:` character.
Typing it will spawn the command bar. To hide it, you can simply delete everything written in it. You will find a complete
list of available commands by typing `:help`.

Here is the description of every element you can find in this world:
* `옷` The player
* `⌴` The boat, needed to go on water. It's too your appearance as soon as you're on water.
* `~` Water, you can go through it as soon as you got boat.
* White blocs : Big trees you cannot cross.
* `K` The key to go deeper in the jungle, by openning a locked door.
* `H` An open door.
* `ℍ` A locked door, requiring a key.

Good luck to find the exit.

## Create your own levels
To create your own levels, you can simply add or edit files in the `levels` directory. The name is corresponding to the level number.
These numbers needs to start at 1, and follow a continuous suite. This file is only a text file, in which you will be able
to put different elements without any problem. Here is all correspondences between characters in text file and in-game elements:
* ` ` (space) Free spaces
* `S` Player spawnpoint
* `B` Boat
* `W` Trees
* `~` Water
* `K` Key
* `L` Locked door
* `O` Open door

When the player ends the last available level, he win the game.

## Future evolutions
The game should evolve, with new enemies, exploiting the already present life system. These enemies's aim will be to
kill the player, or to hold him in the Jungle. When killed, they could loot items, like keys or boats.

A procedural generation can too be added, to have different levels at each gam.

## Credits
Game fully realized by [Yan Buatois](https://github.com/yanbuatois).

Using framework [Termloop](https://github.com/JoelOtter/termloop) by [JoelOtter](https://github.com/JoelOtter)