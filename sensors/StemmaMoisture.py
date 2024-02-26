import struct
import time
from machine import I2C, Pin

class StemmaMoistureSensor():

    def __init__(self):
        self.i2c = I2C(scl=Pin(5), sda=Pin(4), freq=100000)

    def read_moisture(self):
        base = 0x0F
        offset = 0x10

        full_buffer = bytearray([base, offset])
        buf = bytearray(2)

        full_buffer += buf

        self.i2c.writeto(0x36, full_buffer)
        time.sleep(0.005)

        self.i2c.readfrom_into(0x36, buf)
        reading = struct.unpack(">H", buf)[0]

        return reading
    
    def read(self):
        reading = self.read_moisture()

        while reading > 4095:
            reading = self.read_moisture()
        
        return reading




