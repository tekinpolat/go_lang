from datetime import datetime
import requests 

from colorama import Fore,init
init(autoreset=True)

URL = 'https://kur.doviz.com/api/v5/converterItems'

for index in range(10000):
    now = datetime.now()
    response = requests.get(URL)
    print(f'{Fore.GREEN}{now.strftime("%d-%m/%Y, %H:%M:%S")} {index}.sirada {response}')