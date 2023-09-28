export namespace main {
	
	export class UserPreferences {
	    folderPaths: string[];
	    convertTo: string;
	    resizeWidth: number;
	    resizeHeight: number;
	    maintainAspectRatio: boolean;
	    fillUnusedSpace: boolean;
	
	    static createFrom(source: any = {}) {
	        return new UserPreferences(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.folderPaths = source["folderPaths"];
	        this.convertTo = source["convertTo"];
	        this.resizeWidth = source["resizeWidth"];
	        this.resizeHeight = source["resizeHeight"];
	        this.maintainAspectRatio = source["maintainAspectRatio"];
	        this.fillUnusedSpace = source["fillUnusedSpace"];
	    }
	}

}

