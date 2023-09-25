exports.show = (req, res, next) => {
    try {
        const user = [
            {
                "name": "ice"
            }
        ]

        res.status(200).json({
            success: true,
            message: "สำเร็จ",
            data: user,
        });
    } catch (error) {
        res.status(400).json({
            error: {
                message: "เกิดผิดพลาด " + error.message,
            },
        });
    }
};