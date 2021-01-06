#!/usr/bin/env sh

set -e

checkmark="\xe2\x98\x91"
cross="\xe2\x98\x92"

startred='\033[0;31m'
endcolor='\033[0m'
startgreen='\033[0;32m'

echoCheck() {
  echo "$startgreen $checkmark $1 $endcolor"
}

echoError() {
  echo "$startred $cross $1 $endcolor"
}

echob() {
    echo "\033[1m$1\033[0m"
}

check() { # checks if a tool is installed
  if ! type "$1" > /dev/null 2>&1; then
    echoError "$2"
    exit 1
  else
    echoCheck "$3"
  fi
}

goal_prepare() {
  check "multipass" "Multipass is not installed" "Multipass is installed"
}

goal_master() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --disk*)
        __DISK="${1#*=}"
        ;;
      --memory*)
        __MEM="${1#*=}"
        ;;
      --name*)
        __NAME="${1#*=}"
        ;;
      --image*)
        __IMAGE="${1#*=}"
        ;;  
      *)
        printf "***************************\n"
        printf "* Error: Invalid argument.*\n"
        printf "***************************\n"
        exit 1
    esac
    shift
  done

  multipass launch --cloud-init master-cloud-config.yml --verbose --disk "${__DISK:-25G}" --mem "${__MEM:-2G}" --name "${__NAME}" "${IMAGE:-lts}"
}

goal_worker() {
  while [ $# -gt 0 ]; do
    case "$1" in
      --disk*)
        __DISK="${1#*=}"
        ;;
      --memory*)
        __MEM="${1#*=}"
        ;;
      --name*)
        __NAME="${1#*=}"
        ;;
      --image*)
        __IMAGE="${1#*=}"
        ;;  
      *)
        printf "***************************\n"
        printf "* Error: Invalid argument.*\n"
        printf "***************************\n"
        exit 1
    esac
    shift
  done

  multipass launch --cloud-init master-cloud-config.yml --verbose --disk "${__DISK:-25G}" --mem "${__MEM:-2G}" --name "${__NAME}" "${IMAGE:-lts}"
}


#
# Make public functions with goal_ preappended
#
if type "goal_$1" 1> /dev/null; then
  goal_"$1" "${2}"
else
  echo "usage: $0 <goal>

goal:

    prepare         -- check if all tools are installed
    master          -- Install master k8s config on new ubuntu vm
    worker          -- Install worker k8s config on new ubuntu vm

"
  exit 1
fi