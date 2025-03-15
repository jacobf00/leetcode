
def decodeString(s: str) -> str:
    num_st = []
    str_st = []
    curr_num = 0
    curr_str = ""

    for c in s:
        if c.isdigit():
            # Build multi-digit number
            curr_num = curr_num * 10 + int(c)
        elif c == '[':
            # Save the current state
            num_st.append(curr_num)
            str_st.append(curr_str)
            # Reset current values
            curr_num = 0
            curr_str = ""
        elif c == ']':
            # Get previous state
            prev_num = num_st.pop()
            prev_str = str_st.pop()
            # Build new str
            curr_str = prev_str + prev_num * curr_str
        else:
            # Append letter to current str
            curr_str += c

    return curr_str

                


if __name__ == "__main__":
    s = "3[a]2[bc]"
    result = decodeString(s)
    print(result)