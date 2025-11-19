import {expect} from 'chai'
import hre from "hardhat";

const {ethers} = hre;

describe("dichotomy", function () {
    let Dichotomy, dichotomy;
    before(async function () {
        Dichotomy = await ethers.getContractFactory('dichotomy')
        dichotomy = await Dichotomy.deploy();
    })

    it('should test unit', async () => {
        expect(await dichotomy.binarySearch([1, 3, 5, 7, 9], 3)).to.equal(1)
    });
})

// describe("merge", function () {
//     let MergeArray, mergeArray;
//     before(async function () {
//         MergeArray = await ethers.getContractFactory('mergeArray')
//         mergeArray = await MergeArray.deploy();
//     })
//
//     it('should test unit', async () => {
//         expect(await mergeArray.merge([1, 5, 9], [7, 3, 2])).to.deep.equal([1, 2, 3, 5, 7, 9])
//     });
// })

// describe("integer -> roman ", function () {
//     let Roman, roman;
//
//     before(async function () {
//         Roman = await ethers.getContractFactory('roman')
//         roman = await Roman.deploy();
//     })
//
//     it("should convert int to roman", async function () {
//         expect(await roman.integer_to_roman(1994)).to.equals("MCMXCIV")
//         expect(await roman.integer_to_roman(58)).to.equals("LVIII")
//         expect(await roman.integer_to_roman(3749)).to.equals("MMMDCCXLIX")
//     })
//
//     // it("should convert simple numerals", async function () {
//     //     expect(await roman.roman_to_integer("I")).to.equal(1);
//     //     expect(await roman.roman_to_integer("II")).to.equal(2);
//     //     expect(await roman.roman_to_integer("III")).to.equal(3);
//     //     expect(await roman.roman_to_integer("VI")).to.equal(6);   // 5 + 1
//     //     expect(await roman.roman_to_integer("VIII")).to.equal(8); // 5 + 3
//     // });
//     //
//     // it("should convert numerals with subtraction cases", async function () {
//     //     expect(await roman.roman_to_integer("IV")).to.equal(4);  // 5 - 1
//     //     expect(await roman.roman_to_integer("IX")).to.equal(9);  // 10 - 1
//     //     expect(await roman.roman_to_integer("XL")).to.equal(40); // 50 - 10
//     //     expect(await roman.roman_to_integer("XC")).to.equal(90); // 100 - 10
//     //     expect(await roman.roman_to_integer("CD")).to.equal(400); // 500 - 100
//     //     expect(await roman.roman_to_integer("CM")).to.equal(900); // 1000 - 100
//     // });
//     //
//     // it("should convert complex numerals", async function () {
//     //     expect(await roman.roman_to_integer("XIV")).to.equal(14);     // 10 + (5 - 1)
//     //     expect(await roman.roman_to_integer("XLIX")).to.equal(49);   // 40 + 9
//     //     expect(await roman.roman_to_integer("LVIII")).to.equal(58);  // 50 + 5 + 3
//     //     expect(await roman.roman_to_integer("MCMXCIV")).to.equal(1994);
//     // });
//     //
//     // it("should handle long sequences", async function () {
//     //     expect(await roman.roman_to_integer("MMMDCCXXIV")).to.equal(3724);
//     // });
//     //
//     // it("should return 0 for empty input", async function () {
//     //     expect(await roman.roman_to_integer("")).to.equal(0);
//     // });
//
// })
