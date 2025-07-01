GPM, The Unofficial Git Package Manager.
You can do configurations using the dialogue packages based CLI menu configuration's before you build - please note, this is ONLY for pre-build configurations and to change your configurations you would need to rerun the Menu Configuration file and then rebuild and then install the rebuilt Package Manager.

GPM works by Cloning Git (including GitHub or Gitlab repository) repositories - it mainly uses HTTPS to clone but has a SSH feature now.

[NOTE]: To use SSH You will need to have a Public key on your GitHub account - follow the guide given to by GitHub to be able to download Repositories via SSH if you want to.

You can view current configuration's by going into GPM directory and then typing in where the Makefile is "make status" and it should print a table with what Configuration's you had set.

Why use GPM?

- Fast; we use Golang which is a Compiled language which mean's its closer to hardware meaning more peformance.
- Git; it uses Git by Linus Torvalds to Clone repositories.

Just be mindful of what Packages your installing and what Repository it is.
Reasons:
- Git Repo's do not have much Security, what I mean is malware and Malicious Program detection.
- Some Git repositories are NOT the real one, make sure your only using the real repository.
- Some Git Repository's may not be compatible with GPM Autobuild, meaning some will need manual building/configs.

GPM Autobuild:
This software is apart of the gpm-any package (any is any architecture.) that will compile the program.
Some programs may be not compatible with your architecture, Gpm autobuild cannot change architecture code.

It works by checking the Package for a "gpmbuild.json" that contains instructions and dependency lists for GPM.

GPM is primarily a Linux Package Manager.

GPM API:

A addition to GPM, It automate Package installation's via .gapi file instructions.
The GPM-API Programming Interface format is very simple - it is extremely similar to TOML.
Read the GPM-Api READ-ME file for help on Package Automation's.


How to Build:
Just type in make, as long as you have all the Dependency's installed it should work!

Dependency's:
-- Git
-- Python3
-- Go
