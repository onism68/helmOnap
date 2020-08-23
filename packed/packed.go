package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GD4aioSyIAE+Bk4GZLz89Iy0/X10jNLslNTC0JDWBkY+2Xs4hkZGP7/D/Bm5wApZIVqQBilpyeKYpQQwigIpVeSn5sDNi0zwT5+zqSDeVcMBNre13VN/L5rdg6X1u3DXh4xEzISl4rGXOKaufX1jK22E9/mRr0LKFy1g+GRoMUUaSGhPbzep6MX3PAMeHhtjlp83fme//amuxiKZ4hWV7f0JMgvFdFTXtp3VJCRIWJDgOP3M43PDl7Z6TYz3Gm2+rSrJqu+He28UB/q8/rI9MbHal9sdvRmsOtO55dfYBs2aYrzzGW2Xw+V2sokLbjLdc/8RYW39aMrnNyPJlp0W2+9ZPN53ynR+nc2i1NYureqPE26+zV2TeqeEwveXT/frbS8ZQr7GfaHNUHLxWc4e3ZXW0canu7KDPPhtmqTTngeX9AuJxhk56a3yy+38rWdtsnWS7UC0yQv/q6+XNNcU3BvkWjnOpVvJVKMxauzOE58uLt7rS/Dg/M1AmqFa988K7mjcXrjx+RzMYtzsr7P10q802LFZpMS/N3+jEzGtP0vzndffhDRrj5LPlvNVuZSZ7tn8Q1Z2YNffz5/+3T5/Pv34mrr6/8+vbueP357flz83x134mr+1NYJujbr3jR+xhcx906n+J0Wrts3XqzuivwZdOnhZAOflxbSaz+llZfFb1OJXriwngkW5xG9ikGrGRkY2JgRcc7AsIRHHCXO2eBxDo5n1xOO8bAEwwBPMIxMIsyoqY8VKfXBwLZGEIkrLSKbAkp4yM4TQjHFG2EKtmSIMAi7byBAgOG/oz0TA6bfWNlA0swMzAxnGBgYaphAPEAAAAD///X2KSteAwAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
