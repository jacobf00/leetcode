# def closeStrings(word1: str, word2: str) -> bool:
#     word1Set = set(word1)
#     word2Set = set(word2)
#     if word1Set == word2Set:
#         return True
    
#     word1Frequency = {}
#     word2Frequency = {}
#     for c in word1:
#         if c not in word1Frequency:
#             word1Frequency[c] = 1
#         else:
#             word1Frequency[c] += 1
#     for c in word2:
#         if c not in word2Frequency:
#             word2Frequency[c] = 1
#         else:
#             word2Frequency[c] += 1
def closeStrings(word1: str, word2: str) -> bool:
    freq1 = [0] * 26
    freq2 = [0] * 26

    for ch in word1:
        freq1[ord(ch) - ord('a')] += 1

    for ch in word2:
        freq2[ord(ch) - ord('a')] += 1

    for i in range(26):
        if (freq1[i] == 0 and freq2[i] != 0) or (freq1[i] != 0 and freq2[i] == 0):
            return False

    freq1.sort()
    freq2.sort()

    for i in range(26):
        if freq1[i] != freq2[i]:
            return False

    return True


    
    
if __name__ == "__main__":
    word1 = "abc"
    word2 = "bca"

    print(closeStrings(word1,word2))