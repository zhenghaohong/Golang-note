import socket

def main():
    # 1 create tcp
    tcp_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    # 2 create connect
    tcp_socket.connect(('127.0.0.1', 20000))

    # 3 send
    tcp_input = input('Please input you content：')
    tcp_socket.send(tcp_input.encode('utf-8'))  # utf-8

    # 4 close tcp
    tcp_socket.close()


if __name__ == '__main__':
    main()
