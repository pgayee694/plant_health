import time
import StemmaMoisture

sensor = StemmaMoisture.StemmaMoistureSensor()

while True:
    time.sleep(5)
    value = sensor.read()
    print(value)