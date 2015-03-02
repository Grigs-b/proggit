from bs4 import BeautifulSoup
import requests

def get_easter(year):
    data = requests.get('http://www.wheniseastersunday.com/year/{}/'.format(year))
    soup = BeautifulSoup(data.text)
    for p in soup.find_all('p'):
        if "easterdate" in str(p.get("class")):
            return p.text

if __name__ == "__main__":
    year = raw_input("Input Year: ")
    print(get_easter(year))
