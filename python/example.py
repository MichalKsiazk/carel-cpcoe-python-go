import time
from cpcoe import *



cpcoe_device = CPCOE_Device(2, 'COM13')

cpcoe_device.config_univ_channel(9, CPCOE_CHDIR.INPUT, chtype_in=CPCOE_CHTYPE_IN.PT100)


while 1:
    time.sleep(1.0)
    values = cpcoe_device.read_values()
    temp = cpcoe_device.get_univ_ch_val(9)
    print(temp)
    if temp > 27.5:
        cpcoe_device.set_digital_output(6, True, True)
    else:
        cpcoe_device.set_digital_output(6, False, True)