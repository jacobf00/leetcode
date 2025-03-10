# def removeStars(s: str) -> str:
#     before = 0
#     current = 0

#     while current < len(s):
#         if s[current] == "*":
#             s = s[:before] + s[current+1:]
#             current -= 1
#             before -= 1
#         else:
#             if current == 0:
#                 current += 1
#             else:
#                 current += 1
#                 before += 1
#     return s
            
def removeStars(s: str) -> str:
    st = []
    for c in s:
        if c == "*":
            st.pop()
        else:
            st.append(c)
    return "".join(st)


if __name__ == "__main__":
    # s = "leet**cod*e"
    s = "u*ensso****x*q"
    result = removeStars(s)
    print(result)