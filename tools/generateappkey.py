import secrets
import os
from dotenv import load_dotenv, set_key

load_dotenv('../.env')

app_key = secrets.token_hex(32)

set_key('.env', 'APP_KEY', app_key)

print(f"Generated Application Key: {app_key}")