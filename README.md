# Prisoner's Dilemma Game

## Overview

The **Prisoner's Dilemma Game** is a strategic game built with Go, allowing players to explore the classic game theory problem. Players can compete against each other either locally or over a network, choosing between cooperation or defection in each round. The outcome of each decision impacts their "prison time," with the goal of minimizing their own time.

## Features

- **Local and Network Play**: Choose to play against another player on the same machine or over the network.
- **Multiple Rounds**: Customize the number of rounds and view summaries after the game.
- **AI Opponent**: Play against different AI strategies like Tit-for-Tat or Random (future release).
- **Futuristic UI**: Experience a sleek, animated interface with cool ASCII art and a modern design (future release).

## Getting Started

### Prerequisites

- Go 1.20 or higher
- A terminal or command line interface

### Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/kh3rld/prisoners-dilemma.git
    ```
2. **Navigate to the project directory**:
    ```bash
    cd prisoners-dilemma
    ```
3. **Build the project**:
    ```bash
    go build ./cmd
    ```
4. **Run the game**:
    ```bash
    ./cmd
    ```

### Usage

When you start the game, you'll be presented with a menu:

- **1. Play Locally**: Start a local game between two players on the same machine.
- **2. Play Over Network**: Host or join a game over the network.
- **3. View Instructions**: Learn how to play the game.
- **4. Quit**: Exit the game.

Follow the prompts to enter your choices and enjoy the game!

### Playing Over Network

To play over the network:

1. **Host a Game**: Select "Host a Game" on one machine and wait for the other player to join.
2. **Join a Game**: Select "Join a Game" on the second machine and enter the host’s IP address.

### Game Rules

1. Each player can choose to **cooperate** or **defect**.
2. If both players cooperate, they each get **1 year** in prison.
3. If one defects while the other cooperates, the defector goes free and the cooperator gets **3 years** in prison.
4. If both defect, they each get **2 years** in prison.
5. The goal is to minimize your prison time over multiple rounds.

## Contributing

We welcome contributions! Please fork the repository and submit a pull request with your improvements. 

### Future Enhancements

- Add AI opponents with different strategies.
- Implement a leaderboard for networked games.
- Enhance the user interface with more animations and effects.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to the Go community and open-source contributors.
- Inspired by the classic game theory problem, the **Prisoner's Dilemma**.

## Contact

Feel free to reach out via [X](https://x.com/kh3rld) or [GitHub](https://github.com/kh3rld) for any questions or feedback.

