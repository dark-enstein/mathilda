#generate the hashmap of all the reactions and their voltage and save it in a json objecyt

import requests
from bs4 import BeautifulSoup
import json

#parses the html from the website into a dictionary
#initialize arrays and dictionary objects for collecting the results
reaction = []
volt = []
result = {}

#gets the html content from the webpage "http://hyperphysics.phy-astr.gsu.edu/hbase/Tables/electpot.html"
#using beautiful soup
request = requests.get("http://hyperphysics.phy-astr.gsu.edu/hbase/Tables/electpot.html")

soup = BeautifulSoup(request.content, "html.parser")

#using beautiful soup get all the rows of the target table tag that contains the elements of interest (electrochemical cells, and their reactions)
rows = soup.body.table.tr.td.center.table

#iterate through the rows, extracting the test, striping of trailing spaces, and spliting by newline. The result is appended to an array as defined earlier
#Each row consist two columns. Checks are performed to validate if the a specific row is empty, if empty, that iteration is skipped.
for columns in rows:
    if columns == " ":
        continue
    new_col = []
    
    for k in columns:
        new_col.append(k.text.strip().split("\n")[0])
        
    volt.append(new_col[1])
    reaction.append(new_col[0])

#The filledup arrays are appended to the dictionary (result), while defining the title of both columns
result['reaction_title']=reaction[0]
result['volt_title']=volt[0]
del reaction[0]
del volt[0]
result['reactions']=reaction
result['volts']=volt

#The dictionary is then converted into a json object
json_object = json.dumps(result, indent=4 ,ensure_ascii=False) 

#The json object is then saved into a file
with open("seed.json", "w") as outfile:
    outfile.write(json_object)