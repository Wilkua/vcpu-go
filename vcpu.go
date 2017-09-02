package main

import (
	"fmt"
	"log"
	"os"
)

const INS_NOOP = 0x00
const INS_HALT = 0x01
const INS_MOV_IMM_A = 0x02
const INS_MOV_IMM_B = 0x03
const INS_MOV_IMM_C = 0x04
const INS_MOV_IMM_D = 0x05
const INS_MOV_IMM_E = 0x06
const INS_MOV_IMM_F = 0x07
const INS_MOV_MEM_A = 0x08
const INS_MOV_MEM_B = 0x09
const INS_MOV_MEM_C = 0x0A
const INS_MOV_MEM_D = 0x0B
const INS_MOV_MEM_E = 0x0C
const INS_MOV_MEM_F = 0x0D
const INS_MOV_A_MEM = 0x0E
const INS_MOV_B_MEM = 0x0F
const INS_MOV_C_MEM = 0x10
const INS_MOV_D_MEM = 0x11
const INS_MOV_E_MEM = 0x12
const INS_MOV_F_MEM = 0x13
const INS_ADD_IMM = 0x14
const INS_ADD_C_IMM = 0x15
const INS_ADD_MEM = 0x16
const INS_ADD_C_MEM = 0x17
const INS_SUB_IMM = 0x18
const INS_SUB_MEM = 0x19

const DebugEnabled = false

func DebugPrint(s ...interface{}) {
	if DebugEnabled {
		log.Println(s...)
	}
}

func main() {
	var regA byte
	var regB byte
	var regC byte
	var regD byte
	var regE byte
	var regF byte
	var regSS int16
	var regSP int16
	var regFlags int16
	var regPC int16

	var mainMem = make([]byte, 4096) // 4k of RAM

	fmt.Println("Virtual CPU v1.0-alpha")
	fmt.Println("vCPU Instruction Set v01")

	fmt.Println("Reading in program 'prog.bin'...")

	fd, err := os.Open("./prog.bin")
	if err != nil {
		log.Panic(err)
	}

	fd.Read(mainMem)
	fd.Close()

	for {
		var op byte
		var ins = mainMem[regPC]
		regPC++
		DebugPrint("READ:", ins)

		if ins == INS_HALT {
			break
		} else if ins == INS_MOV_IMM_A {
			regA = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_IMM_B {
			regB = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_IMM_C {
			regC = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_IMM_D {
			regD = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_IMM_E {
			regE = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_IMM_F {
			regF = mainMem[regPC]
			regPC++
		} else if ins == INS_MOV_MEM_A {
			op = mainMem[regPC]
			regPC++
			regA = mainMem[op]
		} else if ins == INS_MOV_MEM_B {
			op = mainMem[regPC]
			regPC++
			regB = mainMem[op]
		} else if ins == INS_MOV_MEM_C {
			op = mainMem[regPC]
			regPC++
			regC = mainMem[op]
		} else if ins == INS_MOV_MEM_D {
			op = mainMem[regPC]
			regPC++
			regD = mainMem[op]
		} else if ins == INS_MOV_MEM_E {
			op = mainMem[regPC]
			regPC++
			regE = mainMem[op]
		} else if ins == INS_MOV_MEM_F {
			op = mainMem[regPC]
			regPC++
			regF = mainMem[op]
		} else if ins == INS_MOV_A_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regA
		} else if ins == INS_MOV_B_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regB
		} else if ins == INS_MOV_C_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regC
		} else if ins == INS_MOV_D_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regD
		} else if ins == INS_MOV_E_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regE
		} else if ins == INS_MOV_F_MEM {
			op = mainMem[regPC]
			regPC++
			mainMem[op] = regF
		} else if ins == INS_ADD_IMM {
			nextA := regA + mainMem[regPC]
			regPC++
			if nextA < regA {
				regFlags |= 0x01 // set carry flag
			}
			regA = nextA
		} else if ins == INS_ADD_C_IMM {
			nextA := regA + mainMem[regPC] + byte(regFlags&0x01)
			regPC++
			if nextA < regA {
				regFlags |= 0x01 // set carry flag
			}
			regA = nextA
		} else if ins == INS_ADD_MEM {
			op = mainMem[regPC]
			regPC++
			nextA := regA + mainMem[op]
			if nextA < regA {
				regFlags |= 0x01 // set carry flag
			}
			regA = nextA
		} else if ins == INS_ADD_C_MEM {
			op = mainMem[regPC]
			regPC++
			nextA := regA + mainMem[op] + byte(regFlags&0x01)
			if nextA < regA {
				regFlags |= 0x01 // set carry flag
			}
			regA = nextA
		} else if ins == INS_SUB_IMM {
			nextA := regA - mainMem[regPC]
			regPC++
			if regA < mainMem[regPC] {
				regFlags |= 0x01 // set the carry flag
			}
			regA = nextA
		} else if ins == INS_SUB_MEM {
			op = mainMem[regPC]
			regPC++
			nextA := regA - mainMem[op]
			if regA < mainMem[op] {
				regFlags |= 0x01 // set the carry flag
			}
			regA = nextA
		}
	}

	fmt.Println("\n=== Execution completed ===")
	fmt.Printf("A:  %04X  B:  %04X\n", regA, regB)
	fmt.Printf("C:  %04X  D:  %04X\n", regC, regD)
	fmt.Printf("E:  %04X  F:  %04X\n", regE, regF)
	fmt.Printf("SS: %04X  SP: %04X\n", regSS, regSP)
	fmt.Printf("FLAGS: %016b\n", regFlags)
	fmt.Printf("Program Counter: %04X", regPC)
}
