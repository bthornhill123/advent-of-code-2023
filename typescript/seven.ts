import fs from 'fs';

enum PokerHandScore {
    HighCard = "1",
    OnePair = "2",
    TwoPair = "3",
    ThreeOfAKind = "4",
    FullHouse = "5",
    FourOfAKind = "6",
    FiveOfAKind = "7",
}

export const partOne = () => {
    const cardMap = new Map<string, number>(
        [
            ['A', 14],
            ['K', 13],
            ['Q', 12],
            ['J', 11],
            ['T', 10],
            ['9', 9],
            ['8', 8],
            ['7', 7],
            ['6', 6],
            ['5', 5],
            ['4', 4],
            ['3', 3],
            ['2', 2],
            ['1', 1],
        ]
    )

    const data = fs.readFileSync('seven.txt', 'utf8');
    const lines = data.split('\n');
    const scoredHands: string[] = []
    for (let line of lines) {
        const [hand, bid] = line.split(' ')
        const score = toPokerScore(hand);
        scoredHands.push(`${score}${hand} ${bid}`)
    }

    scoredHands.sort();

    scoredHands.sort((scoreA, scoreB) => {
        for (let i = 0; i < scoreA.length; i++) {
            const charA = cardMap.get(scoreA[i]);
            const charB = cardMap.get(scoreB[i]);
            if (charA === charB) {
                continue;
            }

            if (!charA || !charB) {
                throw new Error('Card not found');
            }

            return charB - charA;
        }
        return 0;
    });

    let totalWinnings = 0;
    let rank = scoredHands.length;
    for (let i = 0; i < scoredHands.length; i++) {
        const bid = scoredHands[i].split(' ')[1];
        totalWinnings += rank * parseInt(bid);
        rank--;
    }

    console.log('totalWinnings :>> ', totalWinnings);
}

const toPokerScore = (hand: string): PokerHandScore => {
    const map = new Map<string, number>()
    for (let letter of hand) {
        const currentValue = map.get(letter) ?? 0
        map.set(letter, currentValue + 1)
    }

    if (map.size === 1) {
        return PokerHandScore.FiveOfAKind
    }

    if (map.size === 2) {
        const values = Array.from(map.values()).sort();
        if (values[0] === 1) {
            return PokerHandScore.FourOfAKind
        } else {
            return PokerHandScore.FullHouse
        }
    }

    if (map.size === 3) {
        const values = Array.from(map.values()).sort();
        if (values[2] === 3) {
            return PokerHandScore.ThreeOfAKind
        } else {
            return PokerHandScore.TwoPair
        }
    }

    if (map.size === 4) {
        return PokerHandScore.OnePair
    }

    return PokerHandScore.HighCard
}

const toPokerScoreWithJokers = (hand: string): PokerHandScore => {
    let jokerCount = 0;

    const map = new Map<string, number>()
    for (let letter of hand) {
        if (letter === 'J') {
            jokerCount++;
            continue;
        }

        const currentValue = map.get(letter) ?? 0
        map.set(letter, currentValue + 1)
    }

    const values = Array.from(map.values()).sort();
    values[values.length - 1] += jokerCount;

    if (values.length === 1) {
        return PokerHandScore.FiveOfAKind
    }

    if (values.length === 2) {
        if (values[0] === 1) {
            return PokerHandScore.FourOfAKind
        } else {
            return PokerHandScore.FullHouse
        }
    }

    if (values.length === 3) {
        if (values[2] === 3) {
            return PokerHandScore.ThreeOfAKind
        } else {
            return PokerHandScore.TwoPair
        }
    }

    if (values.length === 4) {
        return PokerHandScore.OnePair
    }

    return PokerHandScore.HighCard
}

export const partTwo = () => {
    const cardMap = new Map<string, number>(
        [
            ['A', 14],
            ['K', 13],
            ['Q', 12],
            ['T', 10],
            ['9', 9],
            ['8', 8],
            ['7', 7],
            ['6', 6],
            ['5', 5],
            ['4', 4],
            ['3', 3],
            ['2', 2],
            ['1', 1],
            ['J', 0],
        ]
    )

    const data = fs.readFileSync('seven.txt', 'utf8');
    const lines = data.split('\n');
    const scoredHands: string[] = []
    for (let line of lines) {
        const [hand, bid] = line.split(' ')
        const score = toPokerScoreWithJokers(hand);
        scoredHands.push(`${score}${hand} ${bid}`)
    }

    // scoredHands.sort();

    scoredHands.sort((scoreA, scoreB) => {
        for (let i = 0; i < scoreA.length; i++) {
            const charA = cardMap.get(scoreA[i]);
            const charB = cardMap.get(scoreB[i]);
            if (charA === charB) {
                continue;
            }

            if (charA == undefined || charB == undefined) {
                throw new Error('Card not found');
            }

            return charB - charA;
        }
        return 0;
    });

    let totalWinnings = 0;
    let rank = scoredHands.length;
    for (let i = 0; i < scoredHands.length; i++) {
        const bid = scoredHands[i].split(' ')[1];
        totalWinnings += rank * parseInt(bid);
        rank--;
    }

    console.log('totalWinnings :>> ', totalWinnings);
}

// 32T3K -> { }
// T55J5
// KK677
// KTJJT
// QQQJA