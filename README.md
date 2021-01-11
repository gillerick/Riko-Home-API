# Riko Home Automation API
The API provides an interface between the Android application and the ESP32 Module. 
The API receives the transcripted voices from the Android app, decodes them and simplify them (to state and device ID). For instance, if a user gave out the command "Riko, turn ON the bedroom lights", the API translates this to simply [4, 1] where 4 is the ID associated with the bedroom lights. The ESP Module constantly polls for any new command.  
