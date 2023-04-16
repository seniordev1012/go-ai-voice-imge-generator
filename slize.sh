#!/bin/bash

# check the env is Android, Mac, Linux or Windows Then install sound library according to the env
if [ "$OSTYPE" == "linux-android" ]; then
    echo "Android"
    # install sound library
    sudo apt-get install libasound2-dev
elif [ "$OSTYPE" == "darwin"* ]; then
    echo "Mac"
    # install sound library
    brew install portaudio
elif [ "$OSTYPE" == "linux-gnu" ]; then
    echo "Linux"
    # install sound library
    sudo apt-get update
    sudo apt-get install libssl-dev libasound2-dev
    sudo apt-get install libasound2
elif [ "$OSTYPE" == "msys" ]; then
    echo "Windows"
    # install sound library
    pip install pyaudio
else
    echo "Unknown"
fi