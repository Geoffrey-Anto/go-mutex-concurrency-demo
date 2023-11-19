import requests


def test():
    url = "http://localhost:3000/inc"

    while True:
        res1 = requests.post(url, data={})
        print(res1.text)
        res2 = requests.post(url, data={})
        print(res2.text)


test()
