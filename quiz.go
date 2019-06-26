package AminoQuiz

import "fmt"

var (
	ErrInvalidIPInput = fmt.Errorf("err: invalid ip input")
)

const (
	dotNum = 3
)

func calc(nums []int) uint32 {
	return uint32(nums[0]<<24 + nums[1]<<16 + nums[2]<<8 + nums[3])
}

func isAllowed(r uint8) bool {
	if (r >= '0' && r <= '9') || r == ' ' || r == '.' {
		return true
	}
	return false
}

func isDot(r uint8) bool {
	if r == '.' {
		return true
	}
	return false
}

func isSpace(r uint8) bool {
	if r == ' ' {
		return true
	}
	return false
}

func isNum(r uint8) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func isValidIPNum(num int) bool {
	if num >= 0 && num <= 255 {
		return true
	}
	return false
}

func IpConvert(ip string) (uint32, error) {
	var (
		nums       []int
		num        int = -1
		dotSeen    bool
		needDot    bool
		dotCounter int
	)

	for i := 0; i < len(ip); i++ {
		// reject any shit except number, space, dot
		if !isAllowed(ip[i]) {
			return 0, ErrInvalidIPInput
		}

		if isNum(ip[i]) {
			if needDot {
				return 0, ErrInvalidIPInput
			}

			if dotSeen {
				dotSeen = false
			}

			if num == -1 {
				num = 0
			}
			num = num*10 + int(ip[i]-'0')
		}

		if isSpace(ip[i]) {
			// if we get a space, but not dot found yet,
			// then we are in finding dot mode
			if !dotSeen {
				needDot = true
			}
			continue
		}

		if isDot(ip[i]) {
			dotCounter++
			if dotCounter > dotNum {
				return 0, ErrInvalidIPInput
			}

			// every time we get a dot, check if there is a valid number
			if num == -1 || !isValidIPNum(num) {
				return 0, ErrInvalidIPInput
			}

			// hala IP number
			nums = append(nums, num)
			num = -1 // refresh this round of finding number
			needDot = false
			dotSeen = true
		}
	}

	// tailing spaces
	if needDot {
		return 0, ErrInvalidIPInput
	}

	// enough numbers
	if dotCounter != dotNum {
		return 0, ErrInvalidIPInput
	}

	// last number to add
	if num == -1 || !isValidIPNum(num) {
		return 0, ErrInvalidIPInput
	}
	nums = append(nums, num)

	return calc(nums), nil
}
