class Singleton {
    private static instance: Singleton

    private constructor() {
        // 私有构造函数
    }

    public static getInstance(): Singleton {
        if (!Singleton.instance) {
            Singleton.instance = new Singleton()
        }
        return Singleton.instance
    }
}

// test
const s1 = Singleton.getInstance()
const s2 = Singleton.getInstance()
console.log(s1 === s2); // 输出: true
