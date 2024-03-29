package AminoQuiz

import "fmt"

var (
	ErrInvalidIPInput = fmt.Errorf("err: invalid ip input")
)

const (
	dotTotalNum = 3

	ipNumMin = 0
	ipNumMax = 255
)

func IPConvert(ip string) (uint32, error) {
	var (
		nums       [4]int
		dotSeen    bool
		needDot    bool
		dotCounter int

		num = -1 // if -1, means no number fetched from ip string of current round
	)

	for _, r := range ip {
		// reject any shit except number, space, dot
		if !((r >= '0' && r <= '9') || r == ' ' || r == '.') {
			return 0, ErrInvalidIPInput
		}

		if r >= '0' && r <= '9' {
			if needDot {
				return 0, ErrInvalidIPInput
			}

			if dotSeen {
				dotSeen = false
			}

			if num == -1 {
				num = 0
			}
			num = num*10 + int(r-'0')

			// failed ASAP
			if !(num >= ipNumMin && num <= ipNumMax) {
				return 0, ErrInvalidIPInput
			}
		}

		if r == ' ' && !dotSeen {
			// if we get a space, but no dot found yet,
			// then we are in finding dot mode
			needDot = true
		}

		if r == '.' {
			dotCounter++

			// too many dots
			if dotCounter > dotTotalNum {
				return 0, ErrInvalidIPInput
			}

			// every time we get a dot, check if there is a valid number
			if !(num >= ipNumMin && num <= ipNumMax) {
				return 0, ErrInvalidIPInput
			}

			// hala IP number
			nums[dotCounter-1] = num
			num = -1 // refresh this round of finding number
			needDot = false
			dotSeen = true
		}
	}

	// tailing spaces/enough numbers/valid number
	if needDot || dotCounter != dotTotalNum || !(num >= ipNumMin && num <= ipNumMax) {
		return 0, ErrInvalidIPInput
	}

	nums[dotCounter] = num

	return uint32(nums[0]<<24 + nums[1]<<16 + nums[2]<<8 + nums[3]), nil
}
