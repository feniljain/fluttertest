# fluttertest
Custom flutter testing CLI tool for seperate group testing

## Overview
Flutter is a great framework which has helps developers to build fast, responsive UIs at great pace. It is also loaded with an inbuilt testing system. I love testing hence for me VSCode's testing interface was perfect, but a while ago I came across a great editor: VIM. I have been hooked to VIM, and now I use it as my fulltime IDE.

As my love grew for VIM, there were some setbacks which I experienced while using VIM for flutter testing, one can always run "flutter test fileName", for running tests on a file, but this is as far as one could get, it didn't provide me a way to do individual group testing.

This tool helps with just that, just pass in the fileName with file command and you get a list of groups, you can select the group you want and the rest will be handled by the tool.

##USAGE
```bash
Usage:
  fluttertest [flags]
  fluttertest [command]

Available Commands:
  file        Pass the required fileName as the argument
  help        Help about any command

Flags:
  -h, --help   help for fluttertest
```

##CONTRIBUTION
Pull requests are welcome. For major changes, please open an issue first to discuss what would like to change.

##FEATURE REQUESTS
For feature requests one can open an issue and add label "new-feature"

##LICENSE
[MIT](https://choosealicense.com/licenses/mit/)
