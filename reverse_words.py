
def reverseWords(s: str) -> str:
    s = s.strip()
    sList = s.split(" ")
    sList = [w for w in sList if w != ''] 
    sList.reverse()
    return " ".join(sList)

s = "a good   example"
print(reverseWords(s))
