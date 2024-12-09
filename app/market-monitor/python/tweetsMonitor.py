import time
from twitter.scraper import Scraper

def poll_latest_tweets(cookies_file, target_username, poll_interval=10, tweet_limit=5):
    """
    每十秒轮询一次指定用户的最新推文，如果有新推文则打印出来。

    :param cookies_file: 保存 cookies 的文件路径
    :param target_username: 目标用户的用户名
    :param poll_interval: 轮询间隔（秒）
    :param tweet_limit: 获取的推文数量
    """
    # 使用 cookies 文件初始化 Scraper
    scraper = Scraper(cookies=cookies_file)
    
    # 获取目标用户的信息
    users = scraper.users([target_username])
    if not users:
        print(f"未找到用户名为 @{target_username} 的用户。")
        return
    
    target_user = users[0]
    user_id = target_user['id']
    
    last_tweet_id = None
    
    print(f"开始监控用户 @{target_username} 的最新推文，每 {poll_interval} 秒检查一次。")
    
    while True:
        try:
            # 获取目标用户的最新推文
            tweets = scraper.tweets([user_id])
            
            if not tweets:
                print("未获取到推文。")
            else:
                # 如果存在新的推文
                for tweet in reversed(tweets):
                    if last_tweet_id is None or tweet['id'] > last_tweet_id:
                        print(f"新推文ID: {tweet['id']}\n内容: {tweet['text']}\n")
                        last_tweet_id = max(last_tweet_id or 0, tweet['id'])
            
            # 等待下一次轮询
            time.sleep(poll_interval)
        
        except KeyboardInterrupt:
            print("轮询已停止。")
            break
        except Exception as e:
            print(f"发生错误: {e}")
            time.sleep(poll_interval)

if __name__ == "__main__":
    # 替换以下信息为您的 cookies 文件路径和目标用户名
    COOKIES_FILE = "./cookies.json"
    TARGET_USERNAME = "dwours"
    POLL_INTERVAL = 10  # 每10秒轮询一次
    TWEET_LIMIT = 1     # 获取最新的5条推文

    poll_latest_tweets(COOKIES_FILE, TARGET_USERNAME, POLL_INTERVAL, TWEET_LIMIT)