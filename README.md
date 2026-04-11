Script to build the above.

Requires:

https://github.com/swampdawg/go

https://github.com/swampdawg/qtc\#(this repo)

|                     |                            |                            |
| ------------------- | -------------------------- | -------------------------- |
|                     | Prebuilt                   | Sources                    |
| Contents:           | QtCreator 17.0.2           | QtCreator 17.0.2           |
| icu:                | 72.1                       | 72.1                       |
| z3:                 | 4.15.3                     | 4.15.3                     |
| llvm:               | 21.1.0                     | 21.1.0                     |
| cmake:              | 3.30.3                     | 3.30.3                     |
| ninja:              | 1.13.1                     | 1.13.1                     |
| python:             |                            | 3.9.5                      |
| mariadb:            | 11.8                       | 11.8                       |
| qt:                 | 6.9.3                      | 6.9.3                      |
| qtc:                | 17.0.2                     | 17.0.2                     |
| gcc:                | 15.2.0                     | 15.2.0                     |
| node:               | 22.17.0                    | 22.17.0                    |
| md4c:               | 0.4.8                      | 0.4.8                      |
| doxygen:            | 1.9.5                      | 1.9.5                      |
| harfbuzz:           |                            | 11.4.5                     |
| protobuf:           |                            | 3.12.4                     |
| openocd:            | 0.0.0 \#?                  | 0.0.0 \#?                  |
| mosquitto:          | 2.0.15                     | 2.0.15                     |
| qt6ct:              | 0.10                       | 0.10                       |
| git:                | 2.50.0                     | 2.50.0                     |
| libxml2:            |                            | 2.14.5                     |
| distcc:             | 3.4                        | 3.4                        |
| ccache:             |                            | 4.11.3                     |
| binutils:           |                            | 2.43                       |
| newlib:             |                            | 4.4.0.20231231             |
| Pico cross compiler |                            |                            |
| xgcc:               | (pico arm-none-eabi)       | (pico arm-none-eabi)       |
| binutils:           | 2.43                       | 2.43                       |
| gcc:                | 15.2.0                     | 15.2.0                     |
| newlib:             | 4.4.0.20231231             | 4.4.0.20231231             |
| xgcc:(\*1)          | (pico riscv32-unknown-elf) | (pico riscv32-unknown-elf) |
| binutils:           | 2.43                       | 2.43                       |
| gcc:                | 15.2.0                     | 15.2.0                     |
| newlib:             | 4.4.0.20231231             | 4.4.0.20231231             |
| Source              |                            |                            |
| tarball:            | n/a                        | all.tar                    |

\*1) Untested. Author does not own a pico2 at this time.

|                                                                                                                                                                                                                   |                                                        |
| ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ |
| *_link here → _*[*_**Downloa**_*](https://drive.google.com/drive/folders/1F96n3yFwUAq7GNFThEDg-KX7agxr_ZQf)[*_**d**_*](https://drive.google.com/drive/folders/1F96n3yFwUAq7GNFThEDg-KX7agxr_ZQf)** **←link here** |                                                        |
| x86\_64:                                                                                                                                                                                                          | *_qtc-(sdu-linuxmint-21.3-virginia)-17.0-bin.tar.xz'_* |
| md5sum:                                                                                                                                                                                                           | b5de7072bc15ab9addc47cb12943821a                       |
| dev:(\*1)                                                                                                                                                                                                         | qtc-(sdu-linuxmint-21.3-virginia)-17.0.dev             |
| ldd:(\*2)                                                                                                                                                                                                         | qtc-(sdu-linuxmint-21.3-virginia)-17.0.ldd             |
|                                                                                                                                                                                                                   |                                                        |
| aarch64:                                                                                                                                                                                                          | **qtc-(pi24-debian-12-bookworm)-1**7**.0-bin.tar.xz**  |
| md5sum:                                                                                                                                                                                                           | 8dbe89b9381356824b69d8aebab22e16                       |
| dev:(\*1)                                                                                                                                                                                                         | qtc-(pi24-debian-12-bookworm)-17.0.dev                 |
| ldd:(\*2)                                                                                                                                                                                                         | qtc-(pi24-debian-12-bookworm)-17.0.ldd                 |
| Source                                                                                                                                                                                                            |                                                        |
| [*_all.tar_*](https://drive.google.com/file/d/10T9Tmmzg0URicWWA2sXE8uVFRI_bFRRs/view?usp=drive_link)                                                                                                              |                                                        |
| md5sum:                                                                                                                                                                                                           | 554fc1851e892fd7b486a921b38b1d5e                       |

(\*1) This is nothing more than a list of dev packages on the build box.

(\*2) List of QT runtime dependencies. See references to ‘sdldd’ below.

Quick Installation

Binary unpacks to /usr/local/qt/. This can be changed when it is
rebuilt. If you intend to use it as-is, then..

$ mkdir -p \~/usr/src/git/

$ cd usr src/git

$ git clone https://github.com/swampdawg/go

$ git clone https://github.com/swampdawg/qtc

$ mkdir -p \~/bin/

$ cd qtc

$ GOSUB=\~/usr/src/git/go ./go all-install \[path to qtc\*-bin.tar.xz\]

^^^

You should perform the above as the user for which you intend to run
QtCreator.

Invocation:

\#check QT works.. 

$ sd-qt assistant\#wait for indexing to complete then quit 

$ sd-qt designer\#then quit 

Note:

QtCreator stores settings in \~/.config/QtProject/ should you desire to
take steps to preserve settings from an older version. See ‘go qtcrun’
for a hint on how to use “-settingspath”.

Now for QtCreator itself. Point your GUI launcher at..

$ /usr/local/qt/bin/sd-qtc

..or from the command line..

*$ sd-qt sd-qtc*

A suitable icon is at..

/usr/local/qt/share/icons/hicolor/64x64/apps/QtProject-qtcreator.png

The default for the GOSUB variable is “/usr/local/sd/syschk/lib/”. This
isn’t needed for the binary. However it is recommended to..

$ mkdir -p /usr/local/sd/syschk/lib/

$ cd /usr/local/sd/syschk/lib/

..then symlink to every file in “\~/usr/src/git/go/” with the “f\_”
prefix. eg:

f\_go -\> \~/usr/src/git/go/f\_go

f\_tmp -\> \~/usr/src/git/go/f\_tmp

f\_tmpdir -\> \~/usr/src/git/go/f\_tmpdir

f\_trap -\> \~/usr/src/git/go/f\_trap

f\_var -\> \~/usr/src/git/go/f\_var

..at the time of writing. Your life will be a lot easier.

Configure QtCreator

The author likes..

Edit→Preferences→Environment\[Interface\]:Theme=”Flat Dark”.

The fonts may be ugly. Quit QtCreator. Try..

$ sd-qt qt6ct

..ignoring warnings. Relaunch QtCreator.

\~\~\~2do.

Sources

Some sources are verbatim. Some have been generated from github repos.
Some are verbatim but repacked to fit the traditional tarball naming
scheme. 

Dependencies:

\*.dev 

\*.ldd 

As noted earlier, the "dev" list are those packages on the author's
machine at the time. Some may not be relevant to the build. The omission
of a dev package might be important. We are, after all, building
"latest" code on a system with older stuff thus the existence of an
older unwanted system dev package can break things. It depends on the
age gap between your build environment and how new what you’re trying to
build is. Larger the gap, more likely to be issues.

The "ldd" list exists to give an idea of the dependencies to be aiming
for. The ‘sdldd’ tool looked at every qtc-\*bin.tar.xz executable and
dynamic lib and attempted to build a list of all system dependencies.
It’s not bad but isn’t wholly exhaustive. It is also the one item for
which there is no source (this is low down on the “2do” list). For those
who don’t like it, delete it. Its only function is to produce the above
two files.

***Footnotes***

1\) Linux version used is 64Gb mint 21.3 virginia with 1Gb paging.

2\) Raspberry Pi version used is 8Gb rpi5 aarch64 bookworm with 16Gb
paging.

It was possible to use 8Gb rpi4 bullseye but expect an "all-bootstrap"

to take a week. Takes about 12hrs on author's PC. Approx 3days on rpi5.

3\) Most common cause of (seemingly) random build failures is gcc
exhausting

memory, especially when linking collides. Most of the build strives to
use

the clang compiler.

4\) It may be tempting to disable F\_GO\_TMPFS but just set up 16Gb of
paging

before commencing. Your solid state device will thank you. At a push it
is

possible to plug in a thumbdrive and use that as transient swap. Not a

recommended method because it can actually be quite hard to diagnose
when a

thumbdrive fails - they go bad in weird ways.

5\) Pico2 riscv compiler is untested. Author doesn't even have a pico2
atm\!

6\) The usual workflow, when "all-bootstrap" works, is:

a) Use /usr/local/QT/\*/ to build /usr/local/qt/

b) Destroy /usr/local/QT/\*/ and use to B\_QT=/usr/local/qt to start on
next version of /usr/local/QT/\*/

c) When /usr/local/QT/\*/ is verified, backup /usr/local/qt/, wipe it
and rebuild /usr/local/QT/\*/ into /usr/local/qt/.

e) eg:

B\_QT=/usr/local/QT/6500r D\_QT=/usr/local/qt ./go trash-target

B\_QT=/usr/local/QT/6500r D\_QT=/usr/local/qt ./go all-bootstrap

(look in ./BIN/ and rename/archive that elsewhere)

mkdir /usr/local/QT/6700r

(upgrade packages)

B\_QT=/usr/local/qt D\_QT=/usr/local/QT/6700r ./go trash-target

B\_QT=/usr/local/qt D\_QT=/usr/local/QT/6700r ./go all-bootstrap

7\) Installation folders are currently 18Gb each.

8\) The use of "tar -xvJf" to produce .xz is painfully painfully slow
but it

does substantially reduce the size of the binary so we're going to live

with it.

\--------------------------------------------------------------------------------

\~\~\~2do

\--------------------------------------------------------------------------------

1\) Experiment with "install-strip".

2\) Say something about PKG-VER.patch.???

3\) Learn some \*.md stuff\!

^^^scratch that, make a pdf instead\!

4\) picotool

5\) sdldd bug creating \*.ldd file (ie it doesn’t) - string out of
range. fix.
