#!/bin/sh

STATE_OK=0
STATE_ERR=1


while [ $# -gt 0 ];do
    case $1 in
	-f)
	    FILE=$2
	    shift
	    ;;
	*)
	    echo "$1 is not a recognized option!"
	    echo "Wrong input"
	    exit $STATE_UNKNOWN
	    ;; 
    esac
    shift
done

if [ -e "Uploads/$FILE.req" ]; then
    /easy-rsa/easyrsa import-req Uploads/$FILE.req $FILE > /dev/null 2>&1
    echo yes | /easy-rsa/easyrsa sign-req server $FILE > /dev/null 2>&1

    if [ -e "pki/issued/$FILE.crt" ]; then
        exit $STATE_OK
    else
        exit $STATE_ERR
    fi
else
    echo "$FILE does not exist."
    exit $STATE_ERR
fi
exit $STATE_ERR