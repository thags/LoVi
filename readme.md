# Why
I work on multiple applications, each having their own logs scattered throughout the system. I wanted a better way to:
    1. Quickly view the log file I am interested in, which is almost always the most recent one
    2. Tail the log file to see real time logs
    3. Switch between the different log files
    4. Do this on the command line.
    5. Platform independent

Over are the days of opening file explorer, stumbling around searching for the folder containing the logs of the currently working on application. Then, once remembering the folder location, sorting through the files to find the most recent and then opening it in notepad. Now another problem, the file needs to be reloaded everytime there is a log update, which means I need to keep both the notepad open and the file explorer open to the location of the logs, unless I wish to navigate the file explorer once more.

There are other ways to solve some or all of these issue. Initially I was searching for a specific log viewer application, one that will highlight specific lines and automatically tail the file. This did solve somethings, but I do not want to clutter my already full taskbar with another open application. I already have non-negotiable applications open for work, some of my choosing and many forced upon me, but why add yet another? I never want to go to the dreaded multipage taskbar. 

After that I set up a bash script which would do much the same as this application. It was less quick as I still needed to navigate to the folder containing the files and then run the script. I do like this, and with some work it could be just as configurable, but it would always be missing some things. WSL is great, but it as an extra hoop I do not want to jump through to use this, I want it to just work rather that be on windows, linux or mac, all of which I use. There are also feature ideas I have, which I do not think would be feasible to implemnent with bash.

Ultimately, I decided to create something of my own with LoVi, which is just a compressed version of LogViewer.


# Usage
Create a `lovi.config` file in your user home directory

Mac / Linux : `~/`
Windows: `C:\Users\<username>\`

The config file is in json format:
```
{
	"folders":
	[
		{
			"name":"<friendlyName>",
			"filepath":"<Path to directory containing log files>" 
		}
	]
}
```

Call lovi from the command line using `lovi <name from config file>` 
This will open and tail the most recent file in the "filepath" for the given config "name".
