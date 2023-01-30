package human.resource.mgmt.external;

import java.util.Date;
import lombok.Data;

@Data
public class SearchCalendarQuery {

    private Date from;
    private Date to;
    private String userId;
}
