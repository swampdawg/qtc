Script to build the above.

Requires:
https://github.com/swampdawg/go
https://github.com/swampdawg/qtc	#(this repo)
================================================================================
Result:
icu:		72.1
z3:		4.13.0
llvm:		18.1.8
cmake:		3.30.3
ninja:		1.12.1
mariadb:	11.4.2
qt:		6.7.2
qtc:		14.0.1
gcc:		14.2.0
node:		v16.17.0
md4c:		0.4.8
doxygen:	1.9.5
openocd:	0.0.0		#?
mosquitto:	2.0.15
qt6ct:		0.10

xgcc:		(pico arm-none-eabi)
binutils:	2.43
gcc:		14.2.0
newlib:		4.4.0.20231231

xgcc:		(pico riscv32-unknown-elf)	#untested/experimental
binutils:	2.43
gcc:		14.2.0
newlib:		4.4.0.20231231
================================================================================
			Prebuilt version
================================================================================
icu:		72.1
z3:		4.13.0
llvm:		18.1.6
cmake:		3.24.3
ninja:		1.11.1
mariadb:	11.4.2
qt:		6.5
qtc:		11.0
gcc:		14.2.0
node:		v16.17.0
md4c:		0.4.8
doxygen:	1.9.5
openocd:	0.0.0		#?
mosquitto:	2.0.15
qt6ct:		0.10

xgcc:		(pico arm-none-eabi)
binutils:	2.43
gcc:		14.2.0
newlib:		4.4.0.20231231

xgcc:		(pico riscv32-unknown-elf)	#untested/experimental
binutils:	2.43
gcc:		14.2.0
newlib:		4.4.0.20231231

x86_64, linux mint 21.3:	[link]
md5sum:				[value]

aarch64, rpi bookworm:		[link]
md5sum:				[value]

Installation:
$ sudo mkdir /usr/local/QT/
$ sudo chown `id -un`:`id -gn` /usr/local/QT/
$ tar xvJf [tarball]
$ pushd /usr/local/
$ sudo ln -s QT/6500r qt
$ popd
$ cd ~
$ mkdir -p bin
$ pushd bin
$ ln -s /usr/local/QT/6500r/bin/sd-qt
$ popd

Invocation:
#check it works..
$ sd-qt assistant	#wait for indexing to complete then quit
$ sd-qt designer	#then quit

Note:
QtCreator stores settings in ~/.config/QtProject/ should you desire to take
steps to preserve settings from an older version.

$ sd-qt qtcreator &

Use the symlink to refer to the code. eg:
/usr/local/qt/*
..rather than..
/usr/local/QT/6500r/*
..because should you build from source /usr/local/qt/ would be where it would
have been installed.
================================================================================
			Sources
================================================================================
all.tar				[link]
md5sum:				[value]

Some sources are verbatim. Some have been generated from github repos. Some are
verbatim and merely repacked to fit the traditional tarball naming scheme.

Dependencies:
*.dev
*.ldd

The "dev" list are those packages on the author's machine at the time. Some
may not be relevent to the build. The omission of a dev package might be
important. We are, after all, building "latest" code on a system with older
stuff thus the existence of a dev package can break things.

The "ldd" list exists to give an idea of the dependencies to be aiming for.

Instructions:
$ pushd /usr/local/
$ sudo rm -v qt
$ sudo mkdir qt
$ sudo chown `id -un`:`id -gn` qt
$ popd
# install all the dev packages as hinted above.
$ B_QT=/usr/local/QT/6500r D_QT=/usr/local/qt ./go all-bootstrap
#^^^the above takes a very long time and is a waste of time out of the gate.
#^^^only do it once you know the build will work. hint: 'screen' ;-)

More methodically:
Edit 'go' and change..
: ${D_QT:="/usr/local/QT/6500r"}
..to..
: ${D_QT:="/usr/local/qt"}

$ B_QT=/usr/local/QT/6500r ./go bootstrap
#^^^will build a minimal build environment. If this fails you'll be missing
#^^^something fundamental. Start with understanding fcp_bootstrap function.
#^^^look in relevent log/*/*.log file.

eg: 'llvm' fails:
$ B_QT=/usr/local/QT/6500r ./go llvm del obj
#^^^delete OBJ
$ B_QT=/usr/local/QT/6500r ./go llvm cfg
#^^^repeat until prequisites solved.
$ B_QT=/usr/local/QT/6500r ./go llvm mak -j$NPROC
#^^^ditto then manually install..
$ B_QT=/usr/local/QT/6500r ./go llvm ins
$ B_QT=/usr/local/QT/6500r ./go llvm del all
#^^^delete OBJ and SRC
Check it runs to completion..
$ B_QT=/usr/local/QT/6500r ./go llvm all

Once all the phases are complete, check it works from scratch..
$ B_QT=/usr/local/QT/6500r ./go trash-target
$ B_QT=/usr/local/QT/6500r ./go bootstrap

With "bootstrap" working flawlessly it is then possible to move onto..
$ ./go all
..basically repeating the debug approach outlined above. Finally it is
possible to return to..
$ B_QT=/usr/local/QT/6500r ./go all-bootstrap
#^^^only once this succeeds is it possible to remove /usr/local/QT/6500r/ and
#^^^rely on /usr/local/qt/ alone.

================================================================================
			Footnotes
================================================================================
1) Linux version used is 64Gb mint 21.3 virginia with 1Gb paging.
2) Raspberry Pi version used is 8Gb rpi5 aarch64 bookworm with 16Gb paging.
   It was possible to use 8Gb rpi4 bullseye but expect an "all-bootstrap"
   to take a week. Takes about 12hrs on author's PC. Approx 3days on rpi5.
3) Most common cause of (seemingly) random build failures is gcc exhausting
   memory, especially when linking collides. Most of the build strives to use
   the clang compiler.
4) It may be tempting to disable F_GO_TMPFS but just set up 16Gb of paging
   before commencing. Your solid state device will thank you. At a push it is
   possible to plug in a thumbdrive and use that as transient swap. Not a
   recommended method because it can actually be quite hard to diagnose when a
   thumbdrive fails - they go bad in wierd ways.
5) Pico2 riscv compiler is untested. Author doesn't even have a pico2 atm!
6) The usual workflow, when "all-bootstrap" works, is:
   a) Use /usr/local/QT/*/ to build /usr/local/qt/
   b) Destroy /usr/local/QT/*/ and use to B_QT=/usr/local/qt to start on next
      version of /usr/local/QT/*/
   c) When /usr/local/QT/*/ is verified, backup /usr/local/qt/, wipe it and
      rebuild /usr/local/QT/*/ into /usr/local/qt/.
   e) eg:
      B_QT=/usr/local/QT/6500r D_QT=/usr/local/qt ./go trash-target
      B_QT=/usr/local/QT/6500r D_QT=/usr/local/qt ./go all-bootstrap
      (look in ./BIN/ and rename/archive that elsehere)
      mkdir /usr/local/QT/6700r
      (upgrade packages)
      B_QT=/usr/local/qt D_QT=/usr/local/QT/6700r ./go trash-target
      B_QT=/usr/local/qt D_QT=/usr/local/QT/6700r ./go all-bootstrap
7) Installation folders are currently 18Gb each.
8) The use of "tar -xvJf" to produce .xz is painfully painfully slow but it
   does substantially reduce the size of the binary so we're going to live
   with it.
================================================================================
			~~~2do
1) Experiment with "install-strip".
================================================================================
