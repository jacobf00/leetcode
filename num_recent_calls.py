from collections import deque

class RecentCounter:
    # recent time (ms)
    recent_time = 3000

    def __init__(self):
        self.recent_requests = deque()
        

    def ping(self, t: int) -> int:
        self.recent_requests.append(t)
        # recent requests have to be past this time
        earliest_time = t - self.recent_time
        earliest_request = self.recent_requests.popleft()
        # popleft until in the request is in time range
        while earliest_request > 0 and earliest_request < earliest_time:
            earliest_request = self.recent_requests.popleft()
        # Add back the first request that is in the time range
        self.recent_requests.appendleft(earliest_request)
        return len(self.recent_requests)
        
        
        


if __name__ == "__main__":
    times = [[1], [100], [3001], [3002]]
    rc = RecentCounter()

    for t in times:
        recent_requests= rc.ping(t[0])
        print(recent_requests)
    
