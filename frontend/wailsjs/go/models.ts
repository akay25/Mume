export namespace backend {
	
	export class Config {
	    LibraryPath: string;
	    LogLevel: string;
	    WindowWidth: number;
	    WindowHeight: number;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.LibraryPath = source["LibraryPath"];
	        this.LogLevel = source["LogLevel"];
	        this.WindowWidth = source["WindowWidth"];
	        this.WindowHeight = source["WindowHeight"];
	    }
	}

}

