package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GD4aioSyIAE+Bk4GZLz89Iy0/X10jNLslNTC0JDWBkY+2Xs4hkZGP7/D/Bm5wApZIVqQBj1TE4SxSghhFEQSq8kPzcHbNrCbpf4OZPO510xEDj2vq5r4vdds+1ZdU8nqqroeWzZcIlLL7P1zLHSM7fmnSibNc06ZI6WXkMFC9/TthaW/MOXPolplAmF7Jh4+cn088bz3nv3XjssJPbzrqWi8A13xR2+HgudCl0YGMLOhF+w916w8nrpqoAlouG/Niaa8q6qrTt0VV40WOvTVv2NH5m/b3t4d8HXXwdrTY4HTUjSebxFve/2vnbXL1v2NxXr3tq0OHJaqu3WTy2V/ybMb60u/WRSuM4r9ecZK7eTLGHRjl1Xt73fe6trw5kVcmEen22Wtoiwn6nMiw05Lh7h5PjZutpzn8cnnZOb5/g9ZLvxW/rOIvcWkf4NT2qXr/vz3k76ytYQZ4m2TLuKGmOG3qcv9jkeVr7LIZMhxrB5dhaP5/TPt+f6Mjw4N+e5VZGqb7FV7lLNjaWXJ90Mm3XjnXzokRs/rFr2sEh52v4x8V7g+avG4rz/3lTbeSs/RRg+fDDhnMyZW4f/vWBZd97+rZ79vx+/nr26vfv1+fz6ff83Vf999er303ePDq//3K0gles5z6j6wKsXfhEbPoqWbbw4XdT4UeuCx+rrZoYkmkm9/FJT6F+icnPBsv3MsHhPMyviW8/IwMDFjIh3BobcCBmUeGeDxzs4ruueeMTDEg0DPNEwMokwo6ZAVqQUCAPbGkEkrvSIbAoo8SE7TwjFFG+EKdiSIsIg7L6BAAGG/47OTAyYfmNlA0kzMzAznGFgYGhgAvEAAQAA//9zCnEcYgMAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
