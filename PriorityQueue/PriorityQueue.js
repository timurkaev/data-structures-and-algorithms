class PriorityQueue {
    constructor(s) {
        this.maxSize = s;
        this.queArray = [];
        this.nItems = 0;
    }

    insert(item) {
        let j;
        if (this.nItems === 0) {
            this.queArray[this.nItems++] = item;
        } else {
            for (j = this.nItems - 1; j >= 0; j--) {
                if (item > this.queArray[j]) {
                    this.queArray[j + 1] = this.queArray[j];
                } else {
                    break;
                }
            }
            this.queArray[j + 1] = item;
            this.nItems++;
        }
    }

    remove() {
        return this.queArray[--this.nItems];
    }

    peekMin() {
        return this.queArray[this.nItems - 1];
    }

    isEmpty() {
        return this.nItems === 0;
    }

    isFull() {
        return this.nItems === this.maxSize;
    }
}

const thePQ = new PriorityQueue(5);
thePQ.insert(30);
thePQ.insert(50);
thePQ.insert(10);
thePQ.insert(40);
thePQ.insert(20);

while (!thePQ.isEmpty()) {
    const item = thePQ.remove();
    console.log(item);
}
