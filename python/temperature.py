# SPDX-FileCopyrightText: 2021 ladyada for Adafruit Industries
# SPDX-License-Identifier: MIT

import time
import board
import adafruit_dht
import requests
import threading
# Initial the dht device, with data pin connected to:
dhtDevice = adafruit_dht.DHT22(board.D22)

# you can pass DHT22 use_pulseio=False if you wouldn't like to use pulseio.
# This may be necessary on a Linux single board computer like the Raspberry Pi,
# but it will not work in CircuitPython.
# dhtDevice = adafruit_dht.DHT22(board.D18, use_pulseio=False)

external = 0

def fetch_external_temp():
    global external
    while True:
        for _ in range(5):
            try:
                external = float(requests.get("http://192.168.1.241/get_temp", timeout=15).text)
                print("External temp fetched:", external)
                break
            except Exception as error:
                print(error, "when trying to get external temp")
                time.sleep(2.0)
        time.sleep(300)  # Sleep for 5 minutes

thread = threading.Thread(target=fetch_external_temp)
thread.daemon = True
thread.start()
def get_temperature():
    global external
    tot = 0
    N = 3
    real = 0
    for _ in range(N):
        try:
            temperature_c = dhtDevice.temperature
            temperature_f = temperature_c * (9 / 5) + 32
            humidity = dhtDevice.humidity
            print(
                "Temp: {:.1f} F / {:.1f} C    Humidity: {}% ".format(
                    temperature_f, temperature_c, humidity
                )
            )
            tot += temperature_c
            real +=1

        except RuntimeError as error:
            # Errors happen fairly often, DHT's are hard to read, just keep going
            print(error.args[0])
            time.sleep(2.0)
            continue
        except Exception as error:
            dhtDevice.exit()
            raise error

        time.sleep(2.0)



    temp_close =  round(float(tot/real), 2)
    print("external is {}".format(external))
    tot = round( (0.3*temp_close + 0.7*external)  , 2)
    print("tot is {}".format(tot))
    return tot
