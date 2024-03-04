
# Cosmos SDK Blockchain Game: Guess the Number

## Goal and Problem

The goal of this project is to develop a secure and efficient blockchain game using the Cosmos SDK. The game, "Guess the Number," involves a game creator setting a secret number, a reward, and an entry fee. Other players then attempt to guess this number within a specified range and timeframe. The challenge lies in implementing the game logic, ensuring security against vulnerabilities, and providing a practical and user-friendly experience.

## How to Build and Run the Chain

### Get Started

To build and run the blockchain, you can use Ignite CLI:
`ignite chain serve`
This command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com/).

### Web Frontend

Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue` For a React frontend, use: `ignite scaffold react`

These commands can be run within your scaffolded blockchain project. For more information, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Workflow of Each Game

1.  **Game Creation**: A game creator initiates a new game, specifying:
    
    -   Secret number to be guessed.
    -   Maximum time allowed for guessing (duration).
    -   Entry fee (deposit) for submitting a guess.
    -   Reward for correct guesses.
2.  **Guess Submission**: Each player can submit a guess at most 1 time during the game duration.
    
3.  **Game Reveal**: The game creator must reveal the secret number before  `Gameduration+GameExpirationDuration` .
    
4.  **Game Conclusion**:
    
    -   After revealing the number, the game winner and losers are determined.
        
    -   If the game creator does not reveal the number within `Gameduration + GameExpirationDuration` , the chain automatically distributes the prize between all contributors of the game and returns their entry fees.
        
    -   If the game creator  reveal the numbe before  `Gameduration + GameExpirationDuration`, players whose guesses are within the `min_distance_to_win` range share the reward and receive their deposits back. Those who guess incorrectly forfeit their deposits to the game creator.
        

## Commands to Use

### Create a Game

To create a game with the specified parameters, use the following command:
`gtnd tx gtn create-game [number] [duration] [entry-fee] [reward]`
Replace `[number]`, `[duration]`, `[entry-fee]`, and `[reward]` with the appropriate values.

Example:
`gtnd tx gtn create-game 13 12500 10000stake 10stake`

### Submit a Guess

Players can submit their guesses during the game duration using the following command:
`gtnd tx gtn submit-guess [game-id] [number]`
Replace `[game-id]` with the ID of the game and `[number]` with the guessed number.
### Close a Game

To close a game and reveal the secret number, the game creator should use the following command:
`gtnd tx gtn reveal-game [game-id] [salt] [number]`
Replace `[game-id]`, `[salt]`, and `[number]` with the corresponding values.

## Chain Constants

The following chain constants are defined in `x/gtn/types/params.go`:

    Params{
        MaxPlayersPerGame:      100,
        MinDistanceToWin:       100,
        GameExpirationDuration: 1000,
    }

These values can be adjusted as needed.
## Security Considerations

### 1. Secure Storage of Secret Number

To ensure the secrecy of the secret number, the game creator utilizes a random salt and creates a commitment hash. This commitment hash is then stored on-chain. The use of a salt adds randomness and security to the hash, making it computationally infeasible to derive the original secret number from the commitment hash.
More info
https://github.com/Ehsan-saradar/gtn/blob/main/x/gtn/types/keys.go#L40

### 2. Preventing Guess Manipulation

To prevent manipulation of guesses using on-chain data, only the commitment hash of the game is stored on-chain. This ensures that other players cannot use on-chain data and other guesses to increase their chances of winning the game.

### 3. Ensuring Game Creator Integrity

Once the game is created and the commitment hash is stored on-chain, the game creator cannot change the secret number without detection. Any attempt to alter the secret number would result in a different commitment hash, ensuring the integrity of the game creator and the fairness of the game.

## Learn More

-   [Ignite CLI](https://ignite.com/cli)
-   [Tutorials](https://docs.ignite.com/guide)
-   [Ignite CLI Docs](https://docs.ignite.com/)
-   [Cosmos SDK Docs](https://docs.cosmos.network/)
-   [Developer Chat](https://discord.gg/ignite)
