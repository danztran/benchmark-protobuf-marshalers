package just

var engines = []Engine{
	jsonEngine,
	protojsonEngine,
	protoEngine,
	msgpackEngine,
	jsoniterEngine,
	goccyEngine,
	goccyUnorderedEngine,
	goccyFastestEngine,
	sonicStdEngine,
	sonicEngine,
	sonicCompactEngine,
	sonicFastEngine,
}

var samples = []Sample{
	{
		Name:   "Large",
		New:    newProduct,
		Random: genProduct,
	},
	{
		Name:   "Medium",
		New:    newReview,
		Random: genReview,
	},
	{
		Name:   "Small",
		New:    newShipping,
		Random: genShipping,
	},
}
