# venv

add this to any path location
then add this to your .bashrc or .zshrc:

```
venv () {
  /usr/local/bin/venv -x -q -d venv
  . venv/bin/activate
}
```

then anytime you call venv, it will create a python virutalenv in "venv" dir for you

```
-d DIRNAME changes the default dirname [default: venv]
-x creates venv if doesn't exists
-q doesn't display any information on the creation outcome, otherwise will tell you how to activate/deactivate the virtualenv with copy/paste fashion updated paths.
```

**Enjoy**
