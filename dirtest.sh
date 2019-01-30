channelName="testdir"
if ! [ -d /var/${channelName} ]; then
    mkdir -p /var/${channelName}
else
    echo "/var/${channelName} exist!"
fi

