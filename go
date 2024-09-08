#!/bin/bash

NAM=`basename "$0"`
CWD=`pwd`

: ${GOSUB="/usr/local/sd/syschk/lib"}

RETV=

. "$GOSUB""/f_go"

: ${IC_PKG:="icu"}
: ${IC_VER:="72.1"}

: ${Z3_PKG:="z3"}
: ${Z3_VER:="4.13.0"}

: ${LV_PKG:="llvm"}
: ${LV_VER:="18.1.8"}

: ${CM_PKG:="cmake"}
: ${CM_VER:="3.30.3"}

: ${NJ_PKG:="ninja"}
: ${NJ_VER:="1.12.1"}

: ${DB_PKG:="mariadb"}
: ${DB_VER:="11.4.2"}

: ${QT_PKG:="qt"}
: ${QT_VER:="6.7.2"}

: ${QC_PKG:="qtc"}
: ${QC_VER:="14.0.1"}

: ${GC_PKG:="gcc"}
: ${GC_VER:="14.2.0"}

: ${JS_PKG:="node"}
: ${JS_VER:="v16.17.0"}

: ${MD_PKG:="md4c"}
: ${MD_VER:="0.4.8"}

: ${DX_PKG:="doxygen"}
: ${DX_VER:="1.9.5"}

: ${OO_PKG:="openocd"}
: ${OO_VER:="0.0.0"}

: ${MQ_PKG:="mosquitto"}
: ${MQ_VER:="2.0.15"}

: ${C6_PKG:="qt6ct"}
: ${C6_VER:="0.10"}

fcp_rpi5_hack ()
{
 local	r

 r=$( \
	cat /proc/cpuinfo | \
	egrep "^Model[[:space:]]{0,}:" | \
	awk -F':' '{print $2}' \
 )
 case "$r" in
	*Pi*5*)
	echo "$1"",nr_inodes=350k"
	;;

	*)
	echo "$1"
	;;
 esac
}

case `f_go_os_arch` in
	64)
	F_GO_TMPFS=11G
	#F_GO_SAVFS=1
	F_GO_TMPFS=$(fcp_rpi5_hack "$F_GO_TMPFS")
	;;

	*)
	F_GO_TMPFS="11G,nr_inodes=350k"
	;;
esac

PFX=

: ${D_QT:="/usr/local/QT/6700r"}
: ${B_QT:=""}
[ -z "$B_QT" ] && {
CBB="$D_QT""/bin"
CBS="$D_QT""/sbin"
CC="$CBB""/clang"
CXX="$CBB""/clang++"
	} || {
CBB="$B_QT""/bin"
CBS="$B_QT""/sbin"
CC="$CBB""/gcc"
CXX="$CBB""/g++"
}
PTH=$CBS:$CBB":/usr/local/sd/bin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
LDP="$D_QT""/lib64:""$D_QT""/lib:""$D_QT""/lib32"

export PATH="$PTH"
export LD_LIBRARY_PATH="$LDP":$LD_LIBRARY_PATH
export CC CXX

IC_CFG="
--disable-tests
--disable-samples
"

LV_LV_CFG="
-DCMAKE_C_COMPILER=${CBB}/clang
-DCMAKE_CXX_COMPILER=${CBB}/clang++
-DCMAKE_AR=${CBB}/llvm-ar
-DCMAKE_NM=${CBB}/llvm-nm
-DCMAKE_RANLIB=${CBB}/llvm-ranlib
"

LV_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
-DCMAKE_BUILD_TYPE=Release
-DLLVM_ENABLE_RTTI=ON
-DLLVM_PARALLEL_LINK_JOBS=1
-DLLVM_Z3_INSTALL_DIR=${D_QT}
-DLLVM_ENABLE_Z3_SOLVER=ON
-DLLVM_ENABLE_RUNTIMES='all'
-DLLVM_ENABLE_PROJECTS='all'
"

NJ_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
$LV_LV_CFG
"

CM_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
-DCMAKE_PREFIX_PATH=${D_QT}
"
[ -f "$CBB""/clang++" ] && {
CM_CFG="$CM_CFG"" ""$LV_LV_CFG"
}

Z3_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
"
[ -f "$CBB"/bin/z3 ] &&	{
	Z3_CFG="$Z3_CFG"" ""$LV_LV_CFG"
}

DB_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
-DINSTALL_UNIX_ADDRDIR=${D_QT}/lib/mysql.sock
-DBUILD_CONFIG=mysql_release
-DSKIP_TESTS=ON
${LV_LV_CFG}
"

JS_CFG="
--ninja
--prefix=${D_QT}
"

MD_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
${LV_LV_CFG}
"

DX_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
${LV_LV_CFG}
"

QT_CFG="
-release
-nomake examples
-platform sd-linux-clang
-no-pch
-qt-zlib -qt-libjpeg -qt-libpng -qt-pcre -qt-harfbuzz -qt-webp
-prefix ${D_QT}
-opensource
-confirm-license
-icu -I ${D_QT}/include -L ${D_QT}/lib
--
-DCMAKE_ASM_FLAGS="-fno-integrated-as"
"

#QC_CFG="
#CONFIG+=release
#-DCMAKE_INSTALL_PREFIX=${D_QT}
#-DWITH_DOCS=1
#-DWITH_ONLINE_DOCS=1
#-DBUILD_DEVELOPER_DOCS=1
#-DBUILD_WITH_PCH=0
#-DPYTHON_EXECUTABLE=${CBB}/python
#-DPython3_EXECUTABLE=/usr/bin/python3
#${LV_LV_CFG}
#"
QC_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
${LV_LV_CFG}
"

GC_CFG="
--enable-languages=c,c++
"
[ -z "$B_QT" ] || {
GC_CFG="${GC_CFG} --disable-bootstrap"
}

OO_CFG="
--enable-picoprobe
--disable-werror
"

C6_CFG="
-DCMAKE_INSTALL_PREFIX=${D_QT}
"

#dbt (not required for build)
: ${DB_M_BASE:="${D_QT}"}
: ${DB_M_USER:=$(id -un)}
: ${DB_M_PORT:="3306"}
: ${DB_M_BIND:="0.0.0.0"}
: ${DB_M_SOCK:="${D_QT}/lib/mysql.sock"}
: ${DB_R_USER:="root"}
: ${DB_R_PASS:="toor"}
: ${DB_A_USER:="admin"}
: ${DB_A_PASS:="nimda"}

fcp_arc ()
{
 [ -d "$SRC" ] && return 0
 f_go_arc "$@"
}

fcp_cfg ()
{
 f_go_cfg "$@"
}

fcp_ccfg ()
{
 f_go_ccfg "$@"
}

fcp_mak ()
{
 f_go_mak "$@"
}

fcp_cmak ()
{
 f_go_cmak "$@"
}

fcp_ins ()
{
 f_go_ins "$@"
}

fcp_cins ()
{
 f_go_cins "$@"
}

fcp_rem ()
{
 f_go_rem "$@"
}

fcp_del ()
{
 f_go_del "$@"
}

fcp_llvm_cfg ()
{
 [ -x "$CBB/clang" ] || LV_LV_CFG=
 (
 f_go_tmpfs init "$OBJ"
# mkdir -p "$OBJ" || exit 1
 cd "$OBJ" && \
 cmake "$@" $LV_LV_CFG $LV_CFG ../${SRC}/llvm 2>&1 \
 | tee "$CWD""/log/cfg/cfg.""$SRC"".log"
 [ ${PIPESTATUS[0]} -eq 0 ] || return 1
 ) || exit 1
}

fcp_icu_main ()
{
 PKG="$IC_PKG"
 VER="$IC_VER"
 f_go_init
 PFX="$D_QT"
 F_GO_SDIR="/source"

 case "$1" in
        arc)
        shift
        fcp_arc "$@" "$SRC"
        ;;

        cfg)
	shift
	fcp_cfg $IC_CFG "$@"
        ;;

        mak)
        shift
	fcp_mak "$@"
        ;;

        ins)
        shift
	fcp_ins "$@"
        ;;

        rem)
        shift
        fcp_rem "$@"
        ;;

        del)
        shift
        fcp_del "$@"
        ;;

        all)
        fcp_arc -d "$SRC" || exit 1
        fcp_cfg || exit 1
        fcp_mak -j `f_go_bproc` || exit 1
        fcp_ins || exit 1
        fcp_del all
        ;;

        *)
        ;;
 esac
}

fcp_z3_main ()
{
 PKG="$Z3_PKG"
 VER="$Z3_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $Z3_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $Z3_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_llvm_main ()
{
 PKG="$LV_PKG"
 VER="$LV_VER"
 f_go_init

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_llvm_cfg "$@"
	;;

	mak)
	shift
	fcp_cmak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	find "$D_QT" -type f -name 'ompdModule.so' \
		-exec chmod -v u+w '{}' ';'
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_llvm_cfg || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
##
#-type f ! -writeable
#?llvm removes +w on install for this one file?
	find "$D_QT" -type f -name 'ompdModule.so' \
		-exec chmod -v u+w '{}' ';'
##
	;;

	*)
	;;
 esac
}

fcp_cmake_main ()
{
 PKG="$CM_PKG"
 VER="$CM_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $CM_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $CM_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_ninja_main ()
{
 PKG="$NJ_PKG"
 VER="$NJ_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $NJ_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $NJ_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_db_main ()
{
 PKG="$DB_PKG"
 VER="$DB_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $DB_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $DB_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_dbt_mkmysql ()
{
 egrep -q "^""$DB_M_USER"":" /etc/group || {
        groupadd -r "$DB_M_USER" || exit 1
 }
 egrep -q "^""$DB_M_USER"":" /etc/passwd || {
        useradd -r -s /usr/sbin/nologin \
        -g "$DB_M_USER" -d "/usr/local/""$DB_M_USER""/bin" "$DB_M_USER" || exit 1
 }
}

fcp_dbt_init_data ()
{
 (
 cd "$DB_M_BASE"
 ./scripts/mariadb-install-db --no-defaults
 )
}

fcp_dbt_init_user ()
{
 f_tmp_add
cat<<EOF > `f_tmp_top`
CREATE USER '${DB_A_USER}'@'localhost' IDENTIFIED BY '${DB_A_PASS}';
CREATE USER '${DB_A_USER}'@'' IDENTIFIED BY '${DB_A_PASS}';
GRANT ALL PRIVILEGES ON *.* TO '${DB_A_USER}'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO '${DB_A_USER}'@'' WITH GRANT OPTION;
FLUSH PRIVILEGES;
EOF
 sudo "$DB_M_BASE"/bin/mariadb --socket="$DB_M_SOCK" \
	-u "$DB_R_USER" < `f_tmp_top`
 f_tmp_rem
}

fcp_dbt_init_secure ()
{
 local  h=`hostname`

 f_tmp_add
cat<<EOF > `f_tmp_top`
SET PASSWORD FOR '${DB_R_USER}'@'localhost' = PASSWORD('${DB_R_PASS}');
DROP USER ''@'localhost';
DROP USER ''@'${h}';
DROP DATABASE IF EXISTS \`test\`;
FLUSH PRIVILEGES;
EOF
 "$DB_M_BASE""/bin/mariadb" --socket="$DB_M_SOCK" \
	-u "$DB_A_USER" -p"$DB_A_PASS" < `f_tmp_top`
 f_tmp_rem
}

fcp_dbt_1 ()
{
 local	h=`hostname`
 local	p="$D_QT/data/""$h"".pid"

 [ -f "$p" ] && {
	echo "$NAM: -ERR: Already running ($p)?" 1>&2
	exit 1
 }

 "$DB_M_BASE"/bin/mariadbd-safe "$1" --socket="$DB_M_SOCK" --user="$DB_M_USER" &
}

fcp_dbt_0 ()
{
 "$DB_M_BASE"/bin/mariadb-admin --socket="$DB_M_SOCK" \
	-u "$DB_R_USER" -p"$DB_R_PASS" shutdown
}

fcp_dbt_tool ()
{
 case "$1" in
	mkmysql)
	fcp_dbt_mkmysql
	;;

	init-data)
	fcp_dbt_init_data
	;;

	init-user)
	fcp_dbt_init_user
	;;

	init-secure)
	fcp_dbt_init_secure
	;;

	start)
	shift
	fcp_dbt_1 "$@"
	;;

	stop)
	fcp_dbt_0
	;;

	*)
	"$DB_M_BASE"/bin/mariadb --protocol=tcp -P "$DB_M_PORT" \
		-u "$DB_A_USER" -p"$DB_A_PASS" "$@"
	;;
 esac
}

fcp_dbt_mycnf ()
{
 f_tmp_add
cat<<EOF > `f_tmp_top`
[mysqld_safe]

[client]
port		= ${DB_M_PORT}
socket		= ${DB_M_SOCK}

[mysqld]
port		= ${DB_M_PORT}
socket		= ${DB_M_SOCK}
bind-address	= ${DB_M_BIND}
EOF
 cp -vp `f_tmp_top` "$D_QT""/my.cnf"
 f_tmp_rem
}

fcp_dbt_init ()
{
 rm -rf "$D_QT/data"
 fcp_dbt_init_data
 fcp_dbt_1 --no-defaults
 sleep 10
 fcp_dbt_init_user
 fcp_dbt_init_secure
 fcp_dbt_0
 fcp_dbt_mycnf
}

fcp_dbt_dump ()
{
 [ -z "$1" ] && {
	echo "$NAM: fcp_dbt_dump [ sqlfile ]" 1>&2
	exit 1
 }

 "$D_QT""/bin/mariadb-dump" --socket="$DB_M_SOCK" \
	-u "$DB_A_USER" -p"$DB_A_PASS" \
	--lock-tables --all-databases \
	> "$1"
}

fcp_dbt_rest ()
{
 [ -z "$1" ] && {
	echo "$NAM: fcp_dbt_rest [ sqlfile ]" 1>&2
	exit 1
 }
 "$D_QT""/bin/mariadb" --socket="$DB_M_SOCK" \
	-u "$DB_A_USER" -p"$DB_A_PASS" \
	< "$1"
}

# [dbname] { tblname }..
fcp_dbt_save ()
{
 "$D_QT""/bin/mariadb-dump" --socket="$DB_M_SOCK" \
	-u "$DB_A_USER" -p"$DB_A_PASS" \
	--lock-tables \
	"$@"
}

fcp_dbt_load ()
{
 echo "$NAM: Use fcp_dbt_rest!" 1>&2
 exit 1
}

fcp_node_main ()
{
 PKG="$JS_PKG"
 VER="$JS_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	(cd "$SRC" || exit 1
	CC="$CBB/clang -Wno-enum-constexpr-conversion" \
	CXX="$CBB/clang++ -Wno-enum-constexpr-conversion" \
	./configure "$@" $JS_CFG || exit 1
	) || exit 1
	;;

	mak)
	shift
	(cd "$SRC" || exit 1
	make "$@" 2>&1 | tee "$CWD""/log/mak/mak.""$SRC"".log"
	[ ${PIPESTATUS[0]} -eq 0 ] || exit 1
	) || exit 1
	;;

	ins)
	shift
	(cd "$SRC" || exit 1
	make "$@" install || exit 1
	) || exit 1
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_node_main cfg || exit 1
	fcp_node_main mak -j `f_go_bproc` -l `f_go_bproc` || exit 1
	fcp_node_main ins || exit 1
	fcp_del src
	;;

	*)
	;;
 esac
}

fcp_md4c_main ()
{
 PKG="$MD_PKG"
 VER="$MD_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $MD_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $MD_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_doxygen_main ()
{
 PKG="$DX_PKG"
 VER="$DX_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $DX_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $DX_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

#~~~Can't for life of me find answer to why examples (code) will not install
#~~~so we hack it in, probably with some junk but at least qtc [examples]
#~~~box now populated.
fcp_qt_examples ()
{
 case "$1" in
	install)
	find "$SRC" -type d -name examples -exec cp -Rv '{}' "$PFX""/" ';'
	;;

	*)
	;;
 esac
}

fcp_qt_doc ()
{
 (
 (mkdir -p "$OBJ" && mkdir -p "$CWD""/log/doc") || exit 1
 cd "$OBJ" && \
 ninja "$@" 2>&1 | \
 tee "$CWD""/log/doc/doc.""$SRC"".log"
 [ ${PIPESTATUS[0]} -eq 0 ] || return 1
 ) || exit 1

 case "$1" in
	install*)
	fcp_qt_examples install
	;;

	*)
	;;
 esac
}

fcp_qt_arc ()
{
 case "$1" in
	-d)
	fcp_arc "$@"
	[ -d "$SRC""/qtbase/mkspecs/sd-linux-clang" ] || {
		cp -RHv sd-linux-clang "$SRC""/qtbase/mkspecs/" || exit 1
	}
	;;

	*)
	;;
 esac
}

fcp_qt_main ()
{
 PKG="$QT_PKG"
 VER="$QT_VER"
 f_go_init
 PFX=${D_QT}
 export PATH="$OBJ""/qtbase/bin:""$PATH"
 export LD_LIBRARY_PATH="$CWD""/""$OBJ""/qtbase/lib":$LD_LIBRARY_PATH #qvkgen
 export LLVM_INSTALL_DIR=${D_QT}
 export NINJAFLAGS="-v -j"`f_go_bproc`

 case "$1" in
	arc)
	shift
	fcp_qt_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_cfg $QT_CFG "$@"
	;;

	mak)
	shift
	fcp_cmak "$@"
	;;

	ins)
	shift
	fcp_cins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	doc)
	shift
	fcp_qt_doc "$@"
	;;

	all)
	fcp_qt_arc -d "$SRC" || exit 1
	fcp_cfg $QT_CFG || exit 1
	fcp_cmak -j `f_go_bproc` || exit 1
	fcp_cins || exit 1
	[ -f "$CBB/qdoc" ] && {
		fcp_qt_doc docs || exit 1
		fcp_qt_doc install_docs || exit 1
		} || {
		echo "$NAM: No qdoc!" 1>&2
		exit 1
	}
	fcp_del all
	;;

	save)
	f_go_tmpfs save "$SRC"
	f_go_tmpfs save "$OBJ"
	;;

	load)
	f_go_tmpfs load "$SRC"
	f_go_tmpfs load "$OBJ"
	;;

	trash-saves)
	f_go_tmpfs "$@" "$SRC"
	f_go_tmpfs "$@" "$OBJ"
	;;

	*)
	;;
 esac
}

fcp_qtc_main ()
{
 PKG="$QC_PKG"
 VER="$QC_VER"
 f_go_init
 PFX="$D_QT"
 export LLVM_INSTALL_DIR=${D_QT}
 export NINJAFLAGS="-v -j"`f_go_bproc`
 export QTC_ENABLE_CLANG_LIBTOOLING=1
 export Qt5_DIR=${D_QT}
 export INSTALL_ROOT=${D_QT}

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $QC_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $QC_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	f_go_tar
	fcp_del all
	;;

	tar)
	shift
	f_go_tar
	;;

	save)
	f_go_tmpfs save "$SRC"
	f_go_tmpfs save "$OBJ"
	;;

	load)
	f_go_tmpfs load "$SRC"
	f_go_tmpfs load "$OBJ"
	;;

	*)
	;;
 esac
}

fcp_sdqt ()
{
 PKG="$QT_PKG"
 VER="$QT_VER"
 f_go_init
 PFX=${D_QT}

cat<<EOF > "$CWD""/sd-qt"
#!/bin/bash

[ -f $HOME/.picorc ] && . $HOME/.picorc

PATH=${D_QT}/sbin:${D_QT}/bin:\$PATH \\
LD_LIBRARY_PATH=${D_QT}/lib64:${D_QT}/lib:${D_QT}/lib32 \\
LLVM_INSTALL_DIR=${D_QT} \\
QTC_ENABLE_LIBTOOLING=1 \\
QTDIR=${D_QT} \\
MANPATH=${D_QT}/share/man \\
INFOPATH=${D_QT}/share/info \\
"\$@"
EOF

 chmod a+x "$CWD""/sd-qt"
 mkdir -p "$PFX""/bin" || exit 1
 mkdir -p "$PFX""/etc" || exit 1
 install -v "$CWD""/sd-qt" "$PFX""/bin/" || exit 1
 install -v "$CWD""/sd-gdb-multiarch" "$PFX""/bin/" || exit 1
 install -v "$CWD""/gdbinit" "$PFX""/etc/" || exit 1
 install -v "$CWD""/gdbinit.go" "$PFX""/etc/" || exit 1
}

fcp_gcc_req ()
{
 (
 cd "$SRC" || exit 1
 ./contrib/download_prerequisites || exit 1
 ) || exit 1
}

fcp_gcc_arc ()
{
 [ -d "$SRC" ] && return 0
 f_go_arc "$@"
}

fcp_gcc_cfg ()
{
 local CFG

GC_CFG="
$GC_CFG
`f_go_syscc_cfg | egrep "\--(build|host|target)="`
$c
"
 case `f_go_os` in
	raspbian)
GC_CFG="
$GC_CFG
`f_go_syscc_cfg | egrep "\--with-(arch|fpu|float)="`
"
	;;

	*)
	;;
 esac

 CFG=$GC_CFG f_go_cfg "$@"
}

fcp_gcc_main ()
{
 local	mktar=

 PKG="$GC_PKG"
 VER="$GC_VER"
 f_go_init
 PFX="$D_QT"
 REQ="$CWD""/REQ"

 case "$1" in
	arc)
	shift
	fcp_gcc_arc "$@" "$SRC"
	;;

	req)
	fcp_gcc_req "$@"
	;;

	cfg)
	shift
	fcp_gcc_cfg "$@" $GC_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	[ -d "$REQ" ] && {
		find "$REQ" \
		-mindepth 1 -maxdepth 1 -type d -exec rm -rf '{}' ';'
	}
	return 0
	;;

	all)
#	[ -f "$CBB""/gcc" ] || mktar=1
	fcp_gcc_arc -d "$SRC" || exit 1
#	fcp_gcc_req all || exit 1
	fcp_gcc_cfg || exit 1
#	fcp_gcc_cfg $GC_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
#	[ -z "$mktar" ] || f_go_tar || exit 1
	fcp_gcc_main del all
	;;

	qall)
	[ -f `f_go_tar_tarname` ] || {
		echo "$NAM: no \""`f_go_tar_tarname`\"" to extract!" 1>&2
		exit 1
	}
	tar -C / -xvjf `f_go_tar_tarname`
	;;

	*)
	;;
 esac
}

fcp_xgcc_main ()
{
 (
 D_QT="$D_QT"
 CC="$CBB"/gcc CXX="$CBB"/g++
 export D_QT CC CXX
 ./go.xam.none "$@" || exit 1
 ./go.xrv.none "$@" || exit 1
 ) || exit 1
}

fcp_openocd_gen ()
{
 (
 cd "$SRC" || return 1
 ./bootstrap || return 1
 ) || exit 1
}

fcp_openocd_main ()
{
 PKG="$OO_PKG"
 VER="$OO_VER"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	gen)
	shift
	fcp_openocd_gen "$@"
	;;

	cfg)
	shift
	fcp_cfg "$@" $OO_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_openocd_gen || exit 1
	fcp_cfg $OO_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_mqtt_main ()
{
 PKG="$MQ_PKG"
 VER="$MQ_VER"
 f_go_init
 PFX="$D_QT"
CFG="
-DCMAKE_VERBOSE_MAKEFILE=yes
"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $MQ_CFG
	;;

	mak)
	shift
	fcp_cmak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $MQ_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_qt6ct_main ()
{
 PKG="$C6_PKG"
 VER="$C6_VER"
 f_go_init
 PFX="$D_QT"
 export LLVM_INSTALL_DIR=${D_QT}
 export NINJAFLAGS="-v -j"`f_go_bproc`
 export QTC_ENABLE_CLANG_LIBTOOLING=1
 export INSTALL_ROOT=${D_QT}

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@" $C6_CFG
	;;

	mak)
	shift
	fcp_mak "$@"
	;;

	ins)
	shift
	fcp_ins "$@"
	;;

	rem)
	shift
	fcp_rem "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg $C6_CFG || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_picotool_main ()
{
 PKG="picotool"
 VER="2.0.0"
 f_go_init
 PFX="$D_QT"

 case "$1" in
	arc)
	shift
	fcp_arc "$@" "$SRC"
	;;

	cfg)
	shift
	fcp_ccfg "$@"
	;;

	mak)
	shift
	fcp_cmak "$@"
	;;

	ins)
	shift
	fcp_cins "$@"
	;;

	del)
	shift
	fcp_del "$@"
	;;

	all)
	fcp_arc -d "$SRC" || exit 1
	fcp_ccfg || exit 1
	fcp_mak -j `f_go_bproc` || exit 1
	fcp_ins || exit 1
	fcp_del all
	;;

	*)
	;;
 esac
}

fcp_all ()
{
 local	E="./$NAM"
# local	L="sdqt gcc xgcc z3 doxygen llvm cmake ninja db node md4c icu openocd picotool mqtt qt qt6ct qtc"
 local	L="sdqt gcc xgcc z3 doxygen llvm cmake ninja db node md4c icu mqtt qt qt6ct qtc openocd"
 local	i

 [ -z "$PICO_SDK_PATH" ] || L="$L"" picotool"

 for i in $L
 do
	$E $i all || {
		echo "$NAM: Failed on $i" 1>&2
		exit 1
	}
 done

 [ -z "$PICO_SDK_PATH" ] && {
	echo "$NAM: picotool (skipped)" 1>&2
	echo "%NAM: PICO_SDK_PATH not set" 1>&2
 }
}

fcp_ldd ()
{
 local	b

 (
 cd ../QTP/sdldd || exit 1
 D_QT="$D_QT" ./go all || exit 1
  (
  PKG='qtc'
  VER="$QC_VER"
  b=`f_go_tar_tarname`
  b=$(echo "$b" | sed -e 's/-bin\.tar\..*//')
  "$CBB"/sdldd --dev > "$b"".dev" 2>&1
  "$CBB"/sdldd --sp "$D_QT" > "$b"".ldd" 2>&1
  )
 ) || exit 1
}

fcp_bootstrap ()
{
 local	E="./$NAM"
 local	B="sdqt gcc z3 llvm cmake"

 [ -f "$D_QT""/bin/sd-qt" ] && {
	echo "$NAM: $D_QT must be empty!" 2>&1
	echo "$NAM: (hint ./go qtc tar first)" 1>&2
	exit 1
 }

 for i in $B
 do
	$E $i all || {
		echo "$NAM: Failed on $i" 1>&2
		exit 1
	}
 done
}

fcp_trash_target ()
{
 [ -z "$D_QT" ] && {
	echo "$NAM: D_QT empty!" 1>&2
	exit 1
 }
 (
 cd "$D_QT" || exit 1

 find . -mindepth 1 -maxdepth 1 -type d -exec rm -rv '{}' ';'
 find . -mindepth 1 -maxdepth 1 -type f -exec rm -fv '{}' ';'
 )
}

#eg: PKG=qt VER=5.15.2 ./go mkpatch
fcp_mkpatch ()
{
 local	m
 local	o
 local	n=0
 local	f

 f_go_init

 f_tmp_add
 find "$SRC" -type f -name '*.ORIGINAL' > `f_tmp_top`
 while read -r o
 do
	m=`echo "$o" | sed -e 's/\.ORIGINAL//'`
	f="$PKG""-""$VER"".patch""."`printf %03d $n`
	[ -f "$f" ] && {
		echo "$NAM: $f already exists!" 1>&2
		exit 1
	}
	diff -u "$o" "$m" > "$f"
	((n+=1))
 done < `f_tmp_top`
 f_tmp_rem
}

fcp_rsync ()
{
 local	n=`basename "$CWD"`
 local	d=`dirname "$CWD"`

 [ -z "$1" ] && {
	echo "NAME: fcp_rsync [ user@host ]" 1>&2
	exit 1
 }

 (
 cd .. || exit 1
 rsync --progress -auxv \
	--exclude-from="$CWD""/go.rsync" \
	'./'"$n" "$1":"$d""/"
 )
}

fcp_pth ()
{
 local	e="$1"
 shift

 echo P="$PATH"
 echo L="$LD_LIBRARY_PATH"
 echo ~~~

 "$e" "$@"
}

fcp_src ()
{
 local	p v i f d
 local	t="./BIN/rel/src/""$1"".tar"

 [ -z "$1" ] && {
	echo "$NAM: fcp_src [ tarball stem ]" 1>&2
	exit 1
 }
 [ -f "$t" ] && {
	echo "$NAM: fcp_src : $1.tar already exists!" 1>&2
	exit 1
 }

 f_tmp_add
 d=`f_tmp_top`
 f_tmp_add
 egrep "^:.*.._(PKG|VER):" go | awk -F'"' '{print $2}' > `f_tmp_top`
 egrep "^:.*.._(PKG|VER):" go.x??.none | awk -F'"' '{print $2}' >> `f_tmp_top`
 while IFS= read -r p
 do
	read v
	f=$(ls -d "$p"*"$v"*.tar.*)
	ls -d "$p"*"$v"*.patch.??? 2>/dev/null > "$d"
	tar uvhf "$t" "$f" || exit 1
	while IFS= read -r i
	do
		tar uvhf "$t" "$i" || exit 1
	done < "$d"
 done < `f_tmp_top`
 f_tmp_rem
 f_tmp_rem

 mkdir -p `dirname "$t"` || exit 1
 tar uvhf "$t" "$NAM" sd-gdb-multiarch gdbinit gdbinit.go || exit 1
 tar uvhf "$t" "$NAM" go.x??.none || exit 1

 tar uvhf "$t" sd-linux-clang || exit 1

 tar tvf "$t" > `dirname "$t"`"/$1"".txt"
}

fcp_all_install ()
{
 local	t="$1"
 local	s="$HOME""/bin"
 local	q="sd-qt"
 local	c="sd-qtc"
 local	d
 local	r

 [ -f "$t" ] || {
	echo "$NAM: fcp_all_install [ tarball ]" 1>&2
	exit 1
 }

 #assume first line is install folder
 d="/"$(tar tvjf "$t" | head -1 | awk '{print $NF}')"/x"
 d=`dirname "$d"`

 r=$(ls "$d" 2>/dev/null)
 [ -z "$r" ] || {
	echo "$NAM: fcp_all_install Target \"$d\" not empty!" 1>&2
	exit 1
 }

 #install as current user
 sudo chown `id -un`:`id -gn` "$d"
 chmod -v a+x "$d"
 tar -C / -xvjf "$t"

 #set user symlink
 (
 cd "$s" || exit 1
 [ -L "$q" ] && {
	echo "$NAM: fcp_install_all (Retaining old symlink)"
	echo "$NAM: "$(ls -ld "$q")
	} || {
	rm -fv "$q"
 	ln -sv "$d""/bin/""$q"
 }
 ) || {
	echo "$NAM: fcp_all_install ?symlink?" 1>&2
	exit 1
 }

 #make it easier to call qtcreator
 c="$d""/bin/""$c"
 touch "$c"
 chmod -v a+x "$c"
cat<<EOF >"$c"
#!/bin/bash

${d}/bin/${q} qtcreator.sh "\$@"
EOF
 [ -x "$c" ] || {
	echo "$NAM: fcp_all_install ?$c?" 1>&2
	exit 1
 }
 echo "$NAM: Hint: $q assistant"
 echo "$NAM: Hint: $c (<- point gui menu item at this)"
}

f_go_time_b
ARG="$@"
case "$1" in
        icu)
        shift
        fcp_icu_main "$@"
        ;;

	z3)
	shift
	fcp_z3_main "$@"
	;;

	llvm)
	shift
	fcp_llvm_main "$@"
	;;

	cmake)
	shift
	fcp_cmake_main "$@"
	;;

	ninja)
	shift
	fcp_ninja_main "$@"
	;;

	python)
	shift
	fcp_python_main "$@"
	;;

	db)
	shift
	fcp_db_main "$@"
	;;

	node)
	shift
	fcp_node_main "$@"
	;;

	md4c)
	shift
	fcp_md4c_main "$@"
	;;

	doxygen)
	shift
	fcp_doxygen_main "$@"
	;;

	qt)
	shift
	fcp_qt_main "$@"
	;;

	qtc)
	shift
	fcp_qtc_main "$@"
	;;

	gcc)
	shift
	fcp_gcc_main "$@"
	;;

	xgcc)
	shift
	fcp_xgcc_main "$@"
	;;

	sdqt)
	fcp_sdqt
	;;

	openocd)
	shift
	fcp_openocd_main "$@"
	;;

	mqtt)
	shift
	fcp_mqtt_main "$@"
	;;

	qt6ct)
	shift
	fcp_qt6ct_main "$@"
	;;

	all)
	fcp_all
	;;

	bootstrap)
	fcp_bootstrap
	;;

	trash-target)
	fcp_trash_target
	;;

	all-bootstrap)
echo "Bootstrap" >/tmp/"$NAM"
	fcp_bootstrap || exit 1
echo "Stage1" >/tmp/"$NAM"
	B_QT=
	fcp_all || exit 1
echo "Stage2" >/tmp/"$NAM"
	fcp_all || exit 1
echo "Dependencies" >/tmp/"$NAM"
	fcp_ldd || exit 1
	;;

	mkpatch)
	fcp_mkpatch
	;;

	rsync)
	shift
	fcp_rsync "$1"
	;;

	pth)
	shift
	fcp_pth "$@"
	;;

	ldd)
	fcp_ldd
	;;

	cmd)
	#eg /home/foo/usr/src/QT/obj-qt-6.4/qttools/src/linguist
	#eg ../../../../go.006 cmd make
	shift
	"$@" 2>&1 | tee /tmp/cmd.log
	;;

	run)
	shift
	"$D_QT""/bin/sd-qt" "$@"
	;;

	qtcrun)
	rsync -auxv "$HOME/.config/QtProject" /tmp/
	"$D_QT""/bin/sd-qt" qtcreator -settingspath /tmp
	;;

	src)
	shift
	fcp_src "$@"
	;;

	all-install)
	shift
	fcp_all_install "$@"
	;;

	dbt-init)
	shift
	fcp_dbt_init
	;;

	dbt)
	shift
	fcp_dbt_tool "$@"
	;;

	dbt-dump)
	shift
	fcp_dbt_dump "$1"
	;;

	dbt-rest)
	shift
	fcp_dbt_rest "$1"
	;;

	dbt-save)
	shift
	fcp_dbt_save "$@"
	;;

	dbt-load)
	shift
	fcp_dbt_load "$@"
	;;

	dbt-cmd)
	shift
	fcp_dbt_tool -e "$@"
	;;

	picotool)
	shift
	fcp_picotool_main "$@"
	;;

	*)
	;;
esac
RETV=$?
echo "($NAM: CBB=$D_QT)[""$RETV""] (ARG=$ARG)"
f_go_time_e
f_go_time
exit $RETV

##~~~2do:
# all-install needs to ignore 'lost+found'
#
# gcc make install-strip
# cmake --strip
# try install-strip --strip-uneeded -v
##
