def predictPartyVictory(senate: str) -> str:
    banned_dire = 0
    banned_radiant = 0
    
    while True:
        next_round = ""
    
        for sen in senate:
            if sen == 'R':
                if banned_radiant == 0:
                    banned_dire += 1
                    next_round += 'R'
                else:
                    banned_radiant -= 1
            if sen == 'D':
                if banned_dire == 0:
                    banned_radiant += 1
                    next_round += 'D'
                else:
                    banned_dire -= 1
        
        if 'D' not in next_round:
            return "Radiant"
        if 'R' not in next_round:
            return "Dire"

        senate = next_round
    
if __name__ == "__main__":
    # senate = "RD"
    senate = "RDD"
    print(predictPartyVictory(senate))