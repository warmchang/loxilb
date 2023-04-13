#!/bin/bash
source ../common.sh
echo SCENARIO-ulclsctplb
$hexec l3e1 nohup socat -v -T0.5 sctp-l:8080,bind=25.25.25.1,reuseaddr,fork system:"echo 'server1'; cat" >/dev/null 2>&1 &
$hexec l3e2 nohup socat -v -T0.5 sctp-l:8080,bind=26.26.26.1,reuseaddr,fork system:"echo 'server2'; cat" >/dev/null 2>&1 &
$hexec l3e3 nohup socat -v -T0.5 sctp-l:8080,bind=27.27.27.1,reuseaddr,fork system:"echo 'server3'; cat" >/dev/null 2>&1 &

sleep 5
code=0
servArr=( "server1" "server2" "server3" )
ep=( "25.25.25.1" "26.26.26.1" "27.27.27.1" )
ueIp=( "" "32.32.32.1" "31.31.31.1" )
j=0
waitCount=0
while [ $j -le 2 ]
do
    #res=$($hexec ue1 curl ${ep[j]}:8080)
    #res=`$hexec ue1 socat -T10 - SCTP:${ep[j]}:8080,bind=${ueIp[1]}`
    res=`$hexec ue1 socat -T10 - SCTP:${ep[j]}:8080`
    echo $res
    if [[ $res == "${servArr[j]}" ]]
    then
        echo "$res UP"
        j=$(( $j + 1 ))
    else
        echo "Waiting for ${servArr[j]}(${ep[j]})"
        waitCount=$(( $waitCount + 1 ))
        if [[ $waitCount == 10 ]];
        then
            echo "All Servers are not UP"
            echo SCENARIO-ulclsctplb [FAILED]
            #sudo pkill -9 socat 2>&1 > /dev/null
            break
        fi

    fi
    sleep 1
done

for k in {1..2}
do
for i in {1..2}
do
for j in {0..2}
do
    #res=$($hexec ue$k ../common/sctp_client ${ueIp[k]} 88.88.88.88 2020)
    res=$($hexec ue$k socat -T10 - SCTP:88.88.88.88:2020,bind=${ueIp[k]})
    echo -e $res
    if [[ $res != "${servArr[j]}" ]]
    then
        echo -e "Expected ${servArr[j]}, Received : $res"
        if [[ "$res" != *"server"* ]];
        then
            echo "llb1 ct"
            $dexec llb1 loxicmd get ct
            echo "llb2 ct"
            $dexec llb2 loxicmd get ct
            echo "llb2 ip neigh"
            $dexec llb2 ip neigh
        fi
        code=1
    fi
    sleep 1
done
done
done
if [[ $code == 0 ]]
then
    echo SCENARIO-ulclsctplb [OK]
else
    echo SCENARIO-ulclsctplb [FAILED]
fi
sudo pkill -9 socat 2>&1 > /dev/null
exit $code
