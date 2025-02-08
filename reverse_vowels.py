def reverseVowels(s: str) -> str:
    vowels = ['a','e','i','o','u','A','E','I','O','U']
    vowelsInString = {}
    sList = list(s) 
    for i, c in enumerate(s):
        if c in vowels:
            vowelsInString[i] = c

    newKeys = list(vowelsInString.keys())
    newKeys.reverse()
    vowelsInStringReversed = dict(zip(newKeys, vowelsInString.values()))
    for i, c in vowelsInStringReversed.items():
        sList[i] = c
    return ''.join(sList) 

s = "IceCreAm"

print(reverseVowels(s))